package etcd

import "go.etcd.io/etcd/clientv3"

func NewClient(config clientv3.Config) (*clientv3.Client, error) {
	var (
		wc  *clientv3.Client
		err error
	)

	if wc, err = clientv3.New(config); err != nil {
		return nil, err
	}
	return wc, nil
}
