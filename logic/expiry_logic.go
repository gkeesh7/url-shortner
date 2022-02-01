package logic

import (
	"context"
	"log"
	"url-shortner/storage/dao"
)

const _deleteQuery = "DELETE from redirection WHERE expiry < NOW()"

func DeleteExpiredURLs(ctx context.Context) error {
	errDeleting := dao.RedirectionGormImplObj.NativeExec(ctx, _deleteQuery)
	if errDeleting != nil {
		log.Printf("Error while deleting the expired links from DB %v", errDeleting.Error())
		return errDeleting
	}
	return nil
}
