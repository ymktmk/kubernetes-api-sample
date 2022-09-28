package module

// func getPodResult(clientset *kubernetes.Clientset) {
// 	for {
//         namespace := ""
//         name := ""
// 		ctx := context.Background()
//         _, err := clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
//         if errors.IsNotFound(err) {
//             fmt.Printf("Pod %s in namespace %s not found\n", name, namespace)
//         } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
//             fmt.Printf("Error getting pod %s in namespace %s: %v\n",
//                 name, namespace, statusError.ErrStatus.Message)
//         } else if err != nil {
//             panic(err.Error())
//         } else {
//             fmt.Printf("Found pod %s in namespace %s\n", name, namespace)
//         }
//         time.Sleep(10 * time.Second)
//     }
// }
