package v1alpha1

import (
	"reflect"

	datav1alpha1 "github.com/Alluxio/alluxio/api/v1alpha1"
	"github.com/Alluxio/alluxio/pkg/common"
	"github.com/Alluxio/alluxio/pkg/ddc/base"
	"github.com/Alluxio/alluxio/pkg/utils"
)

// Update the dataset with cache status
func (r *RuntimeReconciler) UpdateCacheStates(ctx common.ReconcileRequestContext, engine base.Engine) (dataset *datav1alpha1.Dataset, err error) {
	dataset, err = utils.GetDataset(r.Client, ctx.Name, ctx.Namespace)
	if err != nil {
		r.Log.Error(err, "Failed to get dataset", "name", ctx.Name)
		return dataset, err
	}

	runtime, err := utils.GetRuntime(r.Client, ctx.Name, ctx.Namespace)
	if err != nil {
		r.Log.Error(err, "Failed to get runtime", "name", ctx.Name)
		return dataset, err
	}

	_, err = engine.UpdateRuntimeStatus(runtime)
	if err != nil {
		r.Log.Error(err, "Failed to updatet runtime", "name", ctx.Name)
		return dataset, err
	}

	datasetToUpdate := dataset.DeepCopy()
	datasetToUpdate.Status.CacheStatus.CacheStates = runtime.Status.CacheStates
	// datasetToUpdate.Status.CacheStatus.CacheStates =

	if !reflect.DeepEqual(dataset.Status, datasetToUpdate.Status) {
		err = r.Client.Status().Update(ctx, datasetToUpdate)
		if err != nil {
			r.Log.Error(err, "Update dataset")
			return dataset, err
		}
		dataset = datasetToUpdate
	}

	return dataset, err
}
