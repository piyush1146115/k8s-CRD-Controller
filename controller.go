package main


import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	samplev1alpha1 "github.com/piyush1146115/k8s-CRD-Controller/pkg/apis/statefulcontroller/v1alpha1"
	clientset "github.com/piyush1146115/k8s-CRD-Controller/pkg/generated/clientset/versioned"
	samplescheme "github.com/piyush1146115/k8s-CRD-Controller/pkg/generated/clientset/versioned/scheme"
	informers "github.com/piyush1146115/k8s-CRD-Controller/pkg/generated/informers/externalversions/samplecontroller/v1alpha1"
	listers "github.com/piyush1146115/k8s-CRD-Controller/pkg/generated/listers/samplecontroller/v1alpha1"
)