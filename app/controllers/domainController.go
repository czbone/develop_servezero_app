package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"web/config"
	"web/db"
	"web/modules/log"
	"web/modules/webapp"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"github.com/lithammer/shortuuid"
)

// ######################################################################
// ドメイン情報コントローラ
// ・ドメイン情報を管理する
// ######################################################################
const (
	SITE_CONF_FILE_FORMAT        = "vhost-%s.conf"        // Nginxサイト定義ファイルフォーマット
	BACKUP_SITE_CONF_FILE_FORMAT = "vhost-%s.conf_backup" // Nginxサイト定義ファイルフォーマット
	SITE_CONF_TEMPLATE           = "site.conf.tmpl"       // Nginxサイト定義ファイルテンプレート
	SITE_CONF_PUBLIC_DIR         = "public_html"
	SITE_HOME_DIR_HEAD           = "vhost-"
	MSG_CHANGE_FILENAME          = "ファイル名を変更しました。%s → %s"
)

type DomainController struct{}

type ValidateNewDomain struct {
	Name string `validate:"fqdn"`
}

func (pc *DomainController) Index(c *gin.Context) {
	// パラメータ初期化
	var error, success string // メッセージパラメータ
	domainDb := &db.DomainDb{}

	// 入力値取得
	act := strings.TrimSpace(c.PostForm("act")) // 実行操作

	if act == "add" { // ドメイン追加の場合
		// 入力値取得
		name := strings.TrimSpace(c.PostForm("name")) // ドメイン名

		// 入力値チェック
		newDomain := &ValidateNewDomain{
			Name: name,
		}
		validate := validator.New()
		err := validate.Struct(newDomain)
		if err != nil {
			// エラーメッセージ設定
			for _, err := range err.(validator.ValidationErrors) {
				fieldName := err.Field() // 入力エラーの構造体変数名を取得

				switch fieldName {
				case "Name":
					switch err.Tag() {
					case "fqdn":
						error = "ドメイン名のフォーマットが不正です"
					}
				}
			}
		}

		if error == "" {
			// ドメイン存在確認
			row := domainDb.GetDomainByName(name)
			if row == nil {
				// ドメイン追加
				domainId := generateDomainId(name) // ドメインID作成
				result := domainDb.AddDomain(name, name /*ディレクトリ名*/, domainId)

				// サイト用ディレクトリ作成
				siteDirPath := config.GetEnv().NginxVirtualHostHome + "/" + name
				err = os.MkdirAll(siteDirPath, 0755)
				if err != nil {
					log.Error(err)
				}

				// Webアプリケーションインストール
				webApp, err := webapp.NewWebapp(webapp.WordPressWebAppType)
				if err == nil {
					//webApp.Install(siteDirPath + "/public_html")
					webApp.Install(siteDirPath)
				} else {
					log.Error(err)
				}

				// サイト定義ファイル追加
				installSiteConf(name, domainId)

				// サイト定義格納ディレクトリがある場合はNginxを再起動する
				_, err = os.Stat(config.GetEnv().NginxSiteConfPath)
				if err == nil {
					// Nginxサービスに反映
					restartResult := restartNginxService()
					if !restartResult {
						// 再起動不可の場合は定義ファイルを外して再度実行
						recoverSiteConf(name)

						restartNginxService()
					}
				}

				if result {
					success = "ドメインを登録しました"
				} else {
					error = "ドメイン登録に失敗しました"
				}
			} else {
				error = "登録済みのドメインです"
			}
		}
	} else if act == "del" { // ドメイン削除の場合
		// 入力値取得
		domainId, _ := strconv.Atoi(strings.TrimSpace(c.PostForm("id"))) // ドメインID

		// ドメイン情報取得
		row := domainDb.GetDomain(domainId)
		if row == nil {
			error = "ドメイン削除に失敗しました"
		} else {
			domainName := row["name"]

			// ドメイン削除
			result := domainDb.DelDomain(domainId)
			if result {
				success = fmt.Sprintf("ドメイン「%s」を削除しました", domainName)
			} else {
				error = "ドメイン削除に失敗しました"
			}
		}
	}

	// ドメイン一覧取得
	rows := domainDb.GetDomainList()

	if error == "" { // エラーなしの場合
		// ドメイン一覧表示
		c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
			"app_name":   config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title": "ドメイン一覧",
			"domainList": rows,
			"success":    success,
		})
	} else {
		c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
			"app_name":   config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title": "ドメイン一覧",
			"domainList": rows,
			"error":      error,
		})
	}
}

func generateDomainId(domain string) string {
	domainId := shortuuid.NewWithNamespace(domain + time.Now().Format("2006-01-02 15:04:05"))
	return domainId
}

// Nginxのサイト定義ファイルを作成
func installSiteConf(domain string, domainId string) bool {
	// サイト定義ファイル名生成
	siteConfPath := config.GetEnv().NginxSiteConfPath + "/"
	_, err := os.Stat(siteConfPath)
	if err != nil {
		// デバッグモードの場合はデバッグ用のディレクトリを作成
		if gin.IsDebugging() {
			siteConfPath = "_dest/" + filepath.Base(config.GetEnv().NginxSiteConfPath) + "/"

			// ディレクトリがなければ作成
			os.MkdirAll(siteConfPath, 0755)
		} else {
			log.Error(err)
		}
	}
	siteConfPath += fmt.Sprintf(SITE_CONF_FILE_FORMAT, domain)

	file, err := os.OpenFile(siteConfPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error(err)
		return false
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	siteConfTemplatePath := config.GetEnv().NginxSiteConfTemplateDir + "/" + SITE_CONF_TEMPLATE
	template := pongo2.Must(pongo2.FromFile(siteConfTemplatePath))
	err = template.ExecuteWriter(pongo2.Context{
		"servezero_generated": "ServeZero generate ID: " + domainId,
		"domain_name":         domain,
		"vhost_path":          config.GetEnv().NginxContainerVirtualHostHome + "/" + domain,
		"public_dir":          SITE_CONF_PUBLIC_DIR,
	}, w)
	if err != nil {
		log.Error(err)
		return false
	}
	w.Flush()
	return true
}

// 問題のあるNginxのサイト定義ファイルを退避
func recoverSiteConf(domain string) {
	// サイト定義ファイルを確認
	siteConfPath := config.GetEnv().NginxSiteConfPath + "/" + fmt.Sprintf(SITE_CONF_FILE_FORMAT, domain)
	_, err := os.Stat(siteConfPath)
	if err == nil {
		// ファイル名変更
		backupConfPath := config.GetEnv().NginxSiteConfPath + "/" + fmt.Sprintf(BACKUP_SITE_CONF_FILE_FORMAT, domain)
		err = os.Rename(siteConfPath, backupConfPath)
		if err == nil {
			log.Info(fmt.Sprintf(MSG_CHANGE_FILENAME, filepath.Base(siteConfPath), filepath.Base(backupConfPath)))
		} else {
			log.Error(err)
		}
	}
}

// Nginxにサイト定義を再読み込みさせる
func restartNginxService() bool {
	_, err := exec.Command("docker", "exec", "nginx", "nginx", "-t").Output()
	if err == nil { // テストOKの場合は設定を再読み込み
		exec.Command("docker", "exec", "nginx", "nginx", "-s", "reload").Output()
		return true
	} else {
		log.Error(err)
		return false
	}
}
