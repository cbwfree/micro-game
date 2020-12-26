package app

import (
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
)

// 根据节点ID过滤服务
func FilterNodeId(srvName, id string) selector.Filter {
	return func(old []*registry.Service) []*registry.Service {
		var services []*registry.Service
		var nodeId = fmt.Sprintf("%s-%s", srvName, id)

		for _, service := range old {
			serv := new(registry.Service)
			var nodes []*registry.Node

			for _, node := range service.Nodes {
				if node.Id == nodeId {
					nodes = append(nodes, node)
				}
			}

			// only add service if there's some nodes
			if len(nodes) > 0 {
				// copy
				*serv = *service
				serv.Nodes = nodes
				services = append(services, serv)
			}
		}

		return services
	}
}

// 选择器过滤
func FilterSelector(filter selector.Filter) client.CallOption {
	return client.WithSelectOption(selector.WithFilter(filter))
}

// 选择指定服务节点
func FilterServiceNode(srvName, nodeId string) client.CallOption {
	return FilterSelector(FilterNodeId(srvName, nodeId))
}
