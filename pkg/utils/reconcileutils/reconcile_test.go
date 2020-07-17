package reconcileutils

import (
	"context"

	emperrors "emperror.dev/errors"
	"github.com/golang/mock/gomock"
	"github.com/operator-framework/operator-sdk/pkg/status"
	marketplacev1alpha1 "github.com/redhat-marketplace/redhat-marketplace-operator/pkg/apis/marketplace/v1alpha1"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/utils"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/utils/patch"
	"github.com/redhat-marketplace/redhat-marketplace-operator/test/mock/mock_client"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubectl/pkg/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	flogger "github.com/redhat-marketplace/redhat-marketplace-operator/pkg/utils/logger"
)

var _ = Describe("ReconcileUtils", func() {
	var (
		sut          *testHarness
		ctrl         *gomock.Controller
		client       *mock_client.MockClient
		statusWriter *mock_client.MockStatusWriter
	)

	BeforeEach(func() {
		sut = NewTestHarness()
		ctrl = gomock.NewController(GinkgoT())
		client = mock_client.NewMockClient(ctrl)
		statusWriter = mock_client.NewMockStatusWriter(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	AssertResultsAreStatus := func(status ActionResultStatus) func(result *ExecResult, err error) {
		return func(result *ExecResult, err error) {
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Status).To(Equal(status))
			Expect(result.Err).To(BeNil())
		}
	}

	AssertResultsAreError := func(result *ExecResult, err error) {
		Expect(err).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
		Expect(result.Status).To(Equal(Error))
		Expect(result.Err).To(Equal(sut.testErr))
	}

	It("should return err immediately if get errors", func() {
		gomock.InOrder(
			client.EXPECT().
				Get(sut.ctx, sut.namespacedName, sut.pod).
				Return(sut.testErr).
				Times(1),
		)

		result, err := sut.execClientCommands(client)
		AssertResultsAreError(result, err)
	})

	It("should create if not found", func() {
		gomock.InOrder(
			client.EXPECT().
				Get(sut.ctx, sut.namespacedName, sut.pod).
				Return(errors.NewNotFound(schema.GroupResource{Group: "", Resource: "Pod"}, sut.namespacedName.Name)).
				Times(1),
			client.EXPECT().Create(sut.ctx, sut.pod).Return(nil).Times(1),
			client.EXPECT().Status().Return(statusWriter).Times(1),
			statusWriter.EXPECT().Update(sut.ctx, sut.meterbase).Return(nil).Times(1),
		)

		result, err := sut.execClientCommands(client)
		AssertResultsAreStatus(Requeue)(result, err)
	})

	It("should create and handle error", func() {
		conditions := status.NewConditions(sut.condition)
		sut.meterbase.Status.Conditions = &conditions

		gomock.InOrder(
			client.EXPECT().
				Get(sut.ctx, sut.namespacedName, sut.pod).
				Return(errors.NewNotFound(schema.GroupResource{Group: "", Resource: "Pod"}, sut.namespacedName.Name)).
				Times(1),
			client.EXPECT().Create(sut.ctx, sut.pod).Return(sut.testErr).Times(1),
			client.EXPECT().Status().Return(statusWriter).Times(1),
			statusWriter.EXPECT().Update(sut.ctx, sut.meterbase).Return(nil).Times(1),
		)

		result, err := sut.execClientCommands(client)
		AssertResultsAreError(result, err)
		cond := conditions.GetCondition(marketplacev1alpha1.ConditionError)
		Expect(cond).ToNot(BeNil())
		Expect(cond.Message).To(Equal(sut.testErr.Error()))
	})

	It("should get and update", func() {
		conditions := status.NewConditions(sut.condition)
		sut.meterbase.Status.Conditions = &conditions

		client.EXPECT().Create(sut.ctx, sut.pod).Return(nil).Times(0)
		gomock.InOrder(
			client.EXPECT().
				Get(sut.ctx, sut.namespacedName, sut.pod).
				Return(nil).
				Times(1),
			client.EXPECT().
				Update(sut.ctx, gomock.Any()).
				DoAndReturn(func(ctx context.Context, obj runtime.Object) error {
					if obj == nil {
						Expect(obj).ToNot(BeNil())
					}
					return nil
				}).Times(1),
		)

		result, err := sut.execClientCommands(client)
		AssertResultsAreStatus(Requeue)(result, err)
	})

	It("should get and delete", func() {
		conditions := status.NewConditions(sut.condition)
		sut.meterbase.Status.Conditions = &conditions
		sut.pod.Annotations["foo"] = "bar"

		client.EXPECT().Create(sut.ctx, sut.pod).Return(nil).Times(0)
		client.EXPECT().
			Update(sut.ctx, gomock.Any()).
			Return(nil).Times(0)

		gomock.InOrder(
			client.EXPECT().
				Get(sut.ctx, sut.namespacedName, sut.pod).
				Return(nil).
				Times(1),
			client.EXPECT().
				Delete(sut.ctx, sut.pod).
				Return(nil).
				Times(1),
		)
		result, err := sut.execClientCommands(client)
		AssertResultsAreStatus(Continue)(result, err)
	})
})

type testHarness struct {
	ctx            context.Context
	meterbase      *marketplacev1alpha1.MeterBase
	namespacedName types.NamespacedName
	pod            *corev1.Pod
	updatedPod     *corev1.Pod
	testErr        error
	ignoreNotFound bool
	condition      status.Condition
}

func NewTestHarness() *testHarness {
	harness := &testHarness{}
	harness.meterbase = &marketplacev1alpha1.MeterBase{
		Status: marketplacev1alpha1.MeterBaseStatus{},
	}
	harness.testErr = emperrors.New("a test error")
	harness.namespacedName = types.NamespacedName{Name: "foo", Namespace: "ns"}
	harness.pod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			UID:       types.UID("foo"),
			Name:      "foo",
			Namespace: "bar",
		},
	}
	utils.RhmAnnotator.SetLastAppliedAnnotation(harness.pod)
	harness.ctx = context.TODO()
	harness.condition = status.Condition{
		Type:    marketplacev1alpha1.ConditionInstalling,
		Status:  corev1.ConditionTrue,
		Reason:  marketplacev1alpha1.ReasonMeterBaseStartInstall,
		Message: "created",
	}
	return harness
}

func (h *testHarness) execClientCommands(
	client client.Client,
) (*ExecResult, error) {
	flogger.SetLoggerToZap()
	logger := logf.Log.WithName("clienttest")
	collector := NewCollector()
	patcher := patch.RHMDefaultPatcher

	cc := NewClientCommand(client, scheme.Scheme, logger)
	return cc.Do(
		h.ctx,
		StoreResult(collector.NextPointer("a"), GetAction(h.namespacedName, h.pod)),
		Call(func() (ClientAction, error) {
			getResult := collector.Get("a")
			if getResult.Is(NotFound) {
				return HandleResult(
					StoreResult(collector.NextPointer("b"), CreateAction(
						h.pod,
						CreateWithPatch(utils.RhmAnnotator),
						CreateWithAddOwner(h.pod),
					)),
					OnRequeue(UpdateStatusCondition(h.meterbase, h.meterbase.Status.Conditions, status.Condition{
						Type:    marketplacev1alpha1.ConditionInstalling,
						Status:  corev1.ConditionTrue,
						Reason:  marketplacev1alpha1.ReasonMeterBaseStartInstall,
						Message: "created",
					})),
					OnError(
						Call(func() (ClientAction, error) {
							return UpdateStatusCondition(h.meterbase, h.meterbase.Status.Conditions, status.Condition{
								Type:    marketplacev1alpha1.ConditionError,
								Status:  corev1.ConditionTrue,
								Reason:  marketplacev1alpha1.ReasonMeterBaseStartInstall,
								Message: collector.Get("b").Err.Error(),
							}), nil
						}))), nil
			}

			return nil, nil
		}),
		Call(func() (ClientAction, error) {
			h.updatedPod = h.pod.DeepCopy()
			h.updatedPod.Annotations["foo"] = "bar"

			return HandleResult(
				UpdateWithPatchAction(h.pod, h.updatedPod, patcher),
				OnRequeue(UpdateStatusCondition(h.meterbase, h.meterbase.Status.Conditions, h.condition)),
			), nil
		}),
		HandleResult(
			DeleteAction(h.pod),
			OnNotFound(UpdateStatusCondition(h.meterbase, h.meterbase.Status.Conditions, h.condition)),
		),
	)
}
