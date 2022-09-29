package module

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func ExecController(clientset *kubernetes.Clientset)  {

	watchlist := cache.NewListWatchFromClient(
        clientset.CoreV1().RESTClient(),
        string(v1.ResourcePods),
		// v1.NamespaceAll,
		"default",
        fields.Everything(),
    )

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Note: "NewInformer" はイベント通知を提供しながらストアにデータを入力するためのストアとコントローラーを返します。
	// 機能の低いコントローラー ?
	_, controller := cache.NewInformer(watchlist, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("③ イベントロジック Add")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				// workqueueにenqueueする
				fmt.Println(key) // default/nginx-cd55c47f5-hrgsw
				queue.Add(key)
				fmt.Println(queue.Get()) // default/nginx-cd55c47f5-hrgsw false
			}
			// pod := obj.(*v1.Pod)
			// fmt.Println("Get a pod:", pod.Name, pod.Namespace, pod.Annotations)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("③ イベントロジック Update")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("③ イベントロジック Delete")
		},
	})

	stopCh := make(chan struct{})
    defer close(stopCh)
    go controller.Run(stopCh)
	select {}
   
}
