package module

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

// https://uzimihsr.github.io/post/2020-09-30-kubernetes-client-go-watch-pods/
func Watch(clientset *kubernetes.Clientset)  {

	ctx := context.Background()

	// リソースバージョン
	// resourceVersion := "1"

	// タイムアウト設定
	minWatchTimeout := 5 * time.Minute
	timeoutSeconds := int64(minWatchTimeout.Seconds() * (rand.Float64() + 1.0))
	fmt.Println(timeoutSeconds) // 481

	//　インスタンス生成(監視はいろんな方法がある)
	// kubebuilderはどうなのか？
	w, err := clientset.CoreV1().Events("").Watch(ctx, metav1.ListOptions{
		// 過去のイベントは無視する
		// ResourceVersion: resourceVersion,
		TimeoutSeconds: &timeoutSeconds,
		AllowWatchBookmarks: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	
	// client-goのreflector.goでの実装例(ほぼ省略)
	// Note: タイムアウトになったらブレイクスする処理もあるよ
	for {
		select {
		case event, ok := <- w.ResultChan():
			if !ok {
				break
			}
			switch event.Type {
			case watch.Added:
				// ここで Delta FIFO Queue に入れているっぽい
				fmt.Println("Added: ---------------------------")
				fmt.Println(event)

				meta, err := meta.Accessor(event.Object)
				if err != nil {
					continue
				}
				resourceVersion := meta.GetResourceVersion()
				fmt.Println(resourceVersion)
			case watch.Modified:
				fmt.Println("Modified: ---------------------------")
			case watch.Deleted:
				fmt.Println("Deleted: ---------------------------")
			}
		default:
			// fmt.Println("何もない")
		}	
	}
}
