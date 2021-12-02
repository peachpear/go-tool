package cmd

import (
	"github.com/peachpear/go-tool/internal/compute"
	"github.com/spf13/cobra"
	"log"
)

var expression string

// go run main.go compute -e "1.4 * 2.1"
var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "算数计算",
	Long:  "算数计算功能，实现简易计算",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := compute.Compute(expression)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("结果：%f", content)
	},
}

func init() {
	computeCmd.Flags().StringVarP(&expression, "expression", "e", "", "请输入计算表达式")
}
