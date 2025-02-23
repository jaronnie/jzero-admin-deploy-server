package cmd

import (
	"fmt"
	"net/http"
	"os"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero-contrib/dynamic_conf"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero-admin/server/server/config"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/middleware"
	"github.com/jzero-io/jzero-admin/server/server/handler"
	"github.com/jzero-io/jzero-admin/server/plugins"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:	"server",
	Short:	"server server",
	Long:	"server server",
	Run: func(cmd *cobra.Command, args []string) {
		ss, err := dynamic_conf.NewFsNotify(cfgFile, dynamic_conf.WithUseEnv(true))
		logx.Must(err)
		cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
			Type: "yaml",
		}, ss)
		c, err := cc.GetConfig()
		logx.Must(err)

		// set up logger
		if err := logx.SetUp(c.Log.LogConf); err != nil {
			logx.Must(err)
		}
		if c.Log.LogConf.Mode != "console" {
			logx.AddWriter(logx.NewWriter(os.Stdout))
		}

		printBanner(c)
		fmt.Printf("\nUsing Database: %s\n", c.DatabaseType)
		fmt.Printf("%s conf: %s\n", c.DatabaseType, svc.BuildDataSource(c))
		fmt.Printf("Using Cache: %s\n", c.CacheType)
		logx.Infof("Starting rest server at %s:%d...", c.Rest.Host, c.Rest.Port)

		svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
		run(svcCtx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	server := rest.MustNewServer(svcCtx.MustGetConfig().Rest.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.ErrorCtx(r.Context(), w, err)
	}), rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "X-Request-Id")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}, nil, "*"))
	middleware.Register(server)

	// server add api handlers
	handler.RegisterHandlers(server, svcCtx)

	// server add custom routes
	svcCtx.Custom.AddRoutes(server)

	plugins.LoadPlugins(server, *svcCtx)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(svcCtx.Custom)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
