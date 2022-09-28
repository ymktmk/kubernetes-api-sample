package module

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func ExecController(clientset *kubernetes.Clientset)  {

	watchlist := cache.NewListWatchFromClient(
        clientset.CoreV1().RESTClient(),
        string(v1.ResourcePods),
		// v1.NamespaceAll,
		"default",
        fields.Everything(),
    )

	// Note: "NewInformer" はイベント通知を提供しながらストアにデータを入力するためのストアとコントローラーを返します。
	// 機能の低いコントローラー?
	_, controller := cache.NewInformer(watchlist, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("③ ResourceEventHandler Add --------------------")
			// pod := obj.(*v1.Pod)
			// fmt.Println("Get a pod:", pod.Name, pod.Namespace, pod.Annotations)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("③ ResourceEventHandler Update -----------------")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("③ ResourceEventHandler Delete -----------------")
		},
	})

	stopCh := make(chan struct{})
    defer close(stopCh)
    go controller.Run(stopCh)
	select {}
   
}
