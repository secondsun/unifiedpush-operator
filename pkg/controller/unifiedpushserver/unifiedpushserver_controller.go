package unifiedpushserver

import (
	"context"
	"time"

	pushv1alpha1 "github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1"
	"github.com/aerogear/unifiedpush-operator/pkg/config"
	routev1 "github.com/openshift/api/route/v1"

	openshiftappsv1 "github.com/openshift/api/apps/v1"
	imagev1 "github.com/openshift/api/image/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	cfg = config.New()
	log = logf.Log.WithName("controller_unifiedpushserver")
)

// Add creates a new UnifiedPushServer Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileUnifiedPushServer{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("unifiedpushserver-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &pushv1alpha1.UnifiedPushServer{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource DeploymentConfig and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &openshiftappsv1.DeploymentConfig{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource ImageStream and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &imagev1.ImageStream{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Secret and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource PersistentVolumeClaim and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &corev1.PersistentVolumeClaim{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Service and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource ServiceAccount and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &corev1.ServiceAccount{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Route and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &routev1.Route{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}
	// log.Info("Registering AMQ watches address")
	// //Watch for AMQ Resources
	// err = c.Watch(&source.Kind{Type: &enmassev1beta.Address{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	// })
	// if err != nil {
	// 	return err
	// }
	// log.Info("Registering AMQ watches address space")
	// err = c.Watch(&source.Kind{Type: &enmassev1beta.AddressSpace{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	// })
	// if err != nil {
	// 	return err
	// }

	// log.Info("Registering AMQ watches user")
	// err = c.Watch(&source.Kind{Type: &messaginguserv1beta.MessagingUser{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	// })
	// if err != nil {
	// 	return err
	// }

	// Watch for changes to secondary resource CronJob and requeue the owner UnifiedPushServer
	err = c.Watch(&source.Kind{Type: &batchv1beta1.CronJob{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pushv1alpha1.UnifiedPushServer{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileUnifiedPushServer{}

// ReconcileUnifiedPushServer reconciles a UnifiedPushServer object
type ReconcileUnifiedPushServer struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a UnifiedPushServer object and makes changes based on the state read
// and what is in the UnifiedPushServer.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileUnifiedPushServer) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling UnifiedPushServer")

	// Fetch the UnifiedPushServer instance
	instance := &pushv1alpha1.UnifiedPushServer{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// look for other unifiedPush resources and don't provision a new one if there is another one with Phase=Complete
	existingInstances := &pushv1alpha1.UnifiedPushServerList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "UnifiedPushServer",
			APIVersion: "push.aerogear.org/v1alpha1",
		},
	}
	opts := &client.ListOptions{Namespace: instance.Namespace}
	err = r.client.List(context.TODO(), opts, existingInstances)
	if err != nil {
		reqLogger.Error(err, "Failed to list UnifiedPush resources", "UnifiedPush.Namespace", instance.Namespace)
		return reconcile.Result{}, err
	} else if len(existingInstances.Items) > 1 { // check if > 1 since there's the current one already in that list.
		for _, existingInstance := range existingInstances.Items {
			if existingInstance.Name == instance.Name {
				continue
			}
			if existingInstance.Status.Phase == pushv1alpha1.PhaseProvision || existingInstance.Status.Phase == pushv1alpha1.PhaseComplete {
				reqLogger.Info("There is already a UnifiedPush resource in Complete phase. Doing nothing for this CR.", "UnifiedPush.Namespace", instance.Namespace, "UnifiedPush.Name", instance.Name)
				return reconcile.Result{}, nil
			}
		}
	}

	if instance.Status.Phase == pushv1alpha1.PhaseEmpty {
		instance.Status.Phase = pushv1alpha1.PhaseProvision
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update UnifiedPush resource status phase", "UnifiedPush.Namespace", instance.Namespace, "UnifiedPush.Name", instance.Name)
			return reconcile.Result{}, err
		}
	}
	addressSpaceURL := ""
	//Begin AMQ resource reconcile
	if instance.Spec.UseMessageBroker {
		//check that address space exists
		foundAddressSpace := newAddressSpace(instance)

		// Set UnifiedPushServer instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, foundAddressSpace, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		err = r.client.Get(context.TODO(), types.NamespacedName{Name: foundAddressSpace.Name, Namespace: foundAddressSpace.Namespace}, foundAddressSpace)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Address Space", "AddressSpace.Namespace", foundAddressSpace.Namespace, "AddressSpace.Name", foundAddressSpace.Name)
			err = r.client.Create(context.TODO(), foundAddressSpace)
			if err != nil {
				return reconcile.Result{}, err
			}
			return reconcile.Result{Requeue: true}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		} else if !foundAddressSpace.Status.IsReady {
			return reconcile.Result{Requeue: true}, nil
		} else {
			reqLogger.Info("Found AddressSpace for UPS")
		}

		/*Address space exists and is ready*/

		for _, status := range foundAddressSpace.Status.EndpointStatus {
			if status.Name == "messaging" { //magic value
				addressSpaceURL = status.ServiceHost
			}
		}

		//check that user exists
		foundUser := newMessagingUser(instance)

		// Set UnifiedPushServer instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, foundUser, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		err = r.client.Get(context.TODO(), types.NamespacedName{Name: foundUser.Name, Namespace: foundUser.Namespace}, foundUser)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new MessagingUser", "MessagingUser.Namespace", foundUser.Namespace, "MessagingUser.Name", foundUser.Name)
			err = r.client.Create(context.TODO(), foundUser)
			if err != nil {
				return reconcile.Result{}, err
			}
			return reconcile.Result{Requeue: true}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		} /*User exists and is ready*/

		/*User exists and is ready*/
		//check that addresses exist
		//queues
		queues := []string{"APNsPushMessageQueue", "APNsTokenBatchQueue", "GCMPushMessageQueue", "GCMTokenBatchQueue", "WNSPushMessageQueue", "WNSTokenBatchQueue", "MetricsQueue", "TriggerMetricCollectionQueue", "TriggerVariantMetricCollectionQueue", "BatchLoadedQueue", "AllBatchesLoadedQueue", "FreeServiceSlotQueue"}
		requeueCreate := false
		for _, address := range queues {
			foundQueue := newQueue(instance, address)

			// Set UnifiedPushServer instance as the owner and controller
			if err := controllerutil.SetControllerReference(instance, foundQueue, r.scheme); err != nil {
				return reconcile.Result{}, err
			}

			err = r.client.Get(context.TODO(), types.NamespacedName{Name: foundQueue.Name, Namespace: foundQueue.Namespace}, foundQueue)
			if err != nil && errors.IsNotFound(err) {
				reqLogger.Info("Creating a new Queue", "Queue.Namespace", foundQueue.Namespace, "Queue.Name", foundQueue.Name)
				err = r.client.Create(context.TODO(), foundQueue)
				if err != nil {
					return reconcile.Result{}, err
				}
				requeueCreate = true
			} else if err != nil {
				reqLogger.Info("Queue Error")
				return reconcile.Result{}, err
			} else if !foundQueue.Status.IsReady {
				reqLogger.Info("Queue Not ready", "Queue.Name", foundQueue.Name)
				requeueCreate = true
			}
		}

		if requeueCreate {
			return reconcile.Result{RequeueAfter: time.Second * 5}, nil
		} else {
			reqLogger.Info("Found All queues  for UPS")
		}

		//topics
		topics := []string{"MetricsProcessingStartedTopic", "topic/APNSClient"}
		for _, address := range topics {
			foundTopic := newTopic(instance, address)

			// Set UnifiedPushServer instance as the owner and controller
			if err := controllerutil.SetControllerReference(instance, foundTopic, r.scheme); err != nil {
				return reconcile.Result{}, err
			}

			err = r.client.Get(context.TODO(), types.NamespacedName{Name: foundTopic.Name, Namespace: foundTopic.Namespace}, foundTopic)
			if err != nil && errors.IsNotFound(err) {
				reqLogger.Info("Creating a new Topic", "Topic.Namespace", foundTopic.Namespace, "Topic.Name", foundTopic.Name)
				err = r.client.Create(context.TODO(), foundTopic)
				if err != nil {
					return reconcile.Result{}, err
				}
				requeueCreate = true
			} else if err != nil {
				return reconcile.Result{}, err
			}
		}

		if requeueCreate {
			return reconcile.Result{Requeue: true}, nil
		} else {
			reqLogger.Info("Found All queues and topics for UPS")
		}

	}
	//If AMQ is used, it is ready

	//#region Postgres PVC
	persistentVolumeClaim, err := newPostgresqlPersistentVolumeClaim(instance)
	if err != nil {
		reqLogger.Info("Error making PVC")
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, persistentVolumeClaim, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this PersistentVolumeClaim already exists
	foundPersistentVolumeClaim := &corev1.PersistentVolumeClaim{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: persistentVolumeClaim.Name, Namespace: persistentVolumeClaim.Namespace}, foundPersistentVolumeClaim)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new PersistentVolumeClaim", "PersistentVolumeClaim.Namespace", persistentVolumeClaim.Namespace, "PersistentVolumeClaim.Name", persistentVolumeClaim.Name)
		err = r.client.Create(context.TODO(), persistentVolumeClaim)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion
	//#region Postgres DeploymentConfig
	postgresqlDeploymentConfig, err := newPostgresqlDeploymentConfig(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, postgresqlDeploymentConfig, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this DeploymentConfig already exists
	foundPostgresqlDeploymentConfig := &openshiftappsv1.DeploymentConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: postgresqlDeploymentConfig.Name, Namespace: postgresqlDeploymentConfig.Namespace}, foundPostgresqlDeploymentConfig)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new DeploymentConfig", "DeploymentConfig.Namespace", postgresqlDeploymentConfig.Namespace, "DeploymentConfig.Name", postgresqlDeploymentConfig.Name)
		err = r.client.Create(context.TODO(), postgresqlDeploymentConfig)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region Postgres Service
	postgresqlService, err := newPostgresqlService(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, postgresqlService, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Service already exists
	foundPostgresqlService := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: postgresqlService.Name, Namespace: postgresqlService.Namespace}, foundPostgresqlService)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", postgresqlService.Namespace, "Service.Name", postgresqlService.Name)
		err = r.client.Create(context.TODO(), postgresqlService)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region ServiceAccount
	serviceAccount, err := newUnifiedPushServiceAccount(instance)

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, serviceAccount, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this ServiceAccount already exists
	foundServiceAccount := &corev1.ServiceAccount{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: serviceAccount.Name, Namespace: serviceAccount.Namespace}, foundServiceAccount)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new ServiceAccount", "ServiceAccount.Namespace", serviceAccount.Namespace, "ServiceAccount.Name", serviceAccount.Name)
		err = r.client.Create(context.TODO(), serviceAccount)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region Postgres Secret
	postgresqlSecret, err := newPostgresqlSecret(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, postgresqlSecret, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Secret already exists
	foundPostgresqlSecret := &corev1.Secret{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: postgresqlSecret.Name, Namespace: postgresqlSecret.Namespace}, foundPostgresqlSecret)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Secret", "Secret.Namespace", postgresqlSecret.Namespace, "Secret.Name", postgresqlSecret.Name)
		err = r.client.Create(context.TODO(), postgresqlSecret)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region OauthProxy Service
	oauthProxyService, err := newOauthProxyService(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, oauthProxyService, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Service already exists
	foundOauthProxyService := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: oauthProxyService.Name, Namespace: oauthProxyService.Namespace}, foundOauthProxyService)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", oauthProxyService.Namespace, "Service.Name", oauthProxyService.Name)
		err = r.client.Create(context.TODO(), oauthProxyService)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region UPS Service
	unifiedpushService, err := newUnifiedPushServerService(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, unifiedpushService, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Service already exists
	foundUnifiedpushService := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: unifiedpushService.Name, Namespace: unifiedpushService.Namespace}, foundUnifiedpushService)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", unifiedpushService.Namespace, "Service.Name", unifiedpushService.Name)
		err = r.client.Create(context.TODO(), unifiedpushService)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region OauthProxy Route
	oauthProxyRoute, err := newOauthProxyRoute(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, oauthProxyRoute, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Route already exists
	foundOauthProxyRoute := &routev1.Route{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: oauthProxyRoute.Name, Namespace: oauthProxyRoute.Namespace}, foundOauthProxyRoute)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Route", "Route.Namespace", oauthProxyRoute.Namespace, "Route.Name", oauthProxyRoute.Name)
		err = r.client.Create(context.TODO(), oauthProxyRoute)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region OauthProxy ImageStream
	oauthProxyImageStream, err := newOauthProxyImageStream(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, oauthProxyImageStream, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this ImageStream already exists
	foundOauthProxyImageStream := &imagev1.ImageStream{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: oauthProxyImageStream.Name, Namespace: oauthProxyImageStream.Namespace}, foundOauthProxyImageStream)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new ImageStream", "ImageStream.Namespace", foundOauthProxyImageStream.Namespace, "ImageStream.Name", oauthProxyImageStream.Name)
		err = r.client.Create(context.TODO(), oauthProxyImageStream)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region UPS ImageStream
	unifiedPushImageStream, err := newUnifiedPushImageStream(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Set UnifiedPushServer instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, unifiedPushImageStream, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this ImageStream already exists
	foundUnifiedPushImageStream := &imagev1.ImageStream{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: unifiedPushImageStream.Name, Namespace: unifiedPushImageStream.Namespace}, foundUnifiedPushImageStream)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new ImageStream", "ImageStream.Namespace", unifiedPushImageStream.Namespace, "ImageStream.Name", unifiedPushImageStream.Name)
		err = r.client.Create(context.TODO(), unifiedPushImageStream)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region UPS DeploymentConfig
	unifiedpushDeploymentConfig, err := newUnifiedPushServerDeployment(instance, addressSpaceURL)

	if err := controllerutil.SetControllerReference(instance, unifiedpushDeploymentConfig, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this DeploymentConfig already exists
	foundUnifiedpushDeploymentConfig := &openshiftappsv1.DeploymentConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: unifiedpushDeploymentConfig.Name, Namespace: unifiedpushDeploymentConfig.Namespace}, foundUnifiedpushDeploymentConfig)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new DeploymentConfig", "DeploymentConfig.Namespace", unifiedpushDeploymentConfig.Namespace, "DeploymentConfig.Name", unifiedpushDeploymentConfig.Name)
		err = r.client.Create(context.TODO(), unifiedpushDeploymentConfig)
		if err != nil {
			return reconcile.Result{}, err
		}

		// DeploymentConfig created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}
	//#endregion

	//#region Backups
	if len(instance.Spec.Backups) > 0 {
		backupjobSA := &corev1.ServiceAccount{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "backupjob", Namespace: instance.Namespace}, backupjobSA)
		if err != nil {
			reqLogger.Error(err, "A 'backupjob' ServiceAccount is required for the requested backup CronJob(s). Will check again in 10 seconds")
			return reconcile.Result{RequeueAfter: time.Second * 10}, nil
		}
	}

	existingCronJobs := &batchv1beta1.CronJobList{}
	opts = client.InNamespace(instance.Namespace).MatchingLabels(labels(instance, "backup"))
	err = r.client.List(context.TODO(), opts, existingCronJobs)
	if err != nil {
		return reconcile.Result{}, err
	}

	desiredCronJobs, err := backups(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	for _, desiredCronJob := range desiredCronJobs {
		if err := controllerutil.SetControllerReference(instance, &desiredCronJob, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		if exists := containsCronJob(existingCronJobs.Items, &desiredCronJob); exists {
			err = r.client.Update(context.TODO(), &desiredCronJob)
			if err != nil {
				return reconcile.Result{}, err
			}
		} else {
			reqLogger.Info("Creating a new CronJob", "CronJob.Namespace", desiredCronJob.Namespace, "CronJob.Name", desiredCronJob.Name)
			err = r.client.Create(context.TODO(), &desiredCronJob)
			if err != nil {
				return reconcile.Result{}, err
			}
			return reconcile.Result{}, nil
		}
	}

	for _, existingCronJob := range existingCronJobs.Items {
		desired := containsCronJob(desiredCronJobs, &existingCronJob)
		if !desired {
			reqLogger.Info("Deleting backup CronJob since it was removed from CR", "CronJob.Namespace", existingCronJob.Namespace, "CronJob.Name", existingCronJob.Name)
			err = r.client.Delete(context.TODO(), &existingCronJob)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
	}
	//#endregion

	if foundUnifiedpushDeploymentConfig.Status.ReadyReplicas > 0 && instance.Status.Phase != pushv1alpha1.PhaseComplete {
		instance.Status.Phase = pushv1alpha1.PhaseComplete
		r.client.Status().Update(context.TODO(), instance)
	}

	// Resources already exist - don't requeue
	reqLogger.Info("Skip reconcile: Resources already exist")
	return reconcile.Result{}, nil
}

func containsCronJob(cronJobs []batchv1beta1.CronJob, candidate *batchv1beta1.CronJob) bool {
	for _, cronJob := range cronJobs {
		if candidate.Name == cronJob.Name {
			return true
		}
	}
	return false
}
