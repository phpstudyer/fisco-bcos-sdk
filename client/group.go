/**
 * @Author: ZhaoYadong
 * @Date: 2023-11-18 19:16:13
 * @LastEditors: ZhaoYadong
 * @LastEditTime: 2023-11-20 15:03:28
 * @FilePath: /src/fisco-bcos-sdk/client/group.go
 */

package client

import (
	"context"
	"errors"
	"fmt"
)

type GroupResp struct {
	Code    string
	Message string
	Status  string
}

// GenerateGroup 根据群组ID及创世块参数创建新的群组，本接口仅在兼容性版本为2.2.0及以后的版本有效
func (api *APIHandler) GenerateGroup(ctx context.Context, groupID uint, nodeIds []string, timestamp int64) (string, error) {
	var result GroupResp
	params := map[string]interface{}{
		"timestamp": fmt.Sprintf("%d", timestamp),
		"sealers":   nodeIds,
	}
	err := api.CallContext(ctx, &result, "generateGroup", groupID, params)
	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]
	if result.Code == "0x0" {
		return msg, nil
	}

	return msg, errors.New(msg)
}

// StartGroup 根据群组ID启动相应的群组，本接口仅在兼容性版本为2.2.0及以后的版本有效
func (api *APIHandler) StartGroup(ctx context.Context, groupID uint) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "startGroup", groupID)
	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]

	if result.Code == "0x0" {
		return msg, nil
	}

	return msg, errors.New(GroupCodeMessage[result.Code])
}

// StopGroup 根据群组ID停止相应的群组，本接口仅在兼容性版本为2.2.0及以后的版本有效
func (api *APIHandler) StopGroup(ctx context.Context, groupID uint) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "stopGroup", groupID)
	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]

	if result.Code == "0x0" {
		return msg, nil
	}

	return msg, errors.New(GroupCodeMessage[result.Code])
}

// RemoveGroup 根据群组ID删除相应群组，群组数据会被保留以供将来恢复群组，本接口仅在兼容性版本为2.2.0及以后的版本有效
func (api *APIHandler) RemoveGroup(ctx context.Context, groupID uint) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "removeGroup", groupID)

	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]

	if result.Code == "0x0" {
		return msg, nil
	}

	return msg, errors.New(GroupCodeMessage[result.Code])
}

// RecoverGroup 根据群组ID恢复相应群组，本接口仅在兼容性版本为2.2.0及以后的版本有效
func (api *APIHandler) RecoverGroup(ctx context.Context, groupID uint) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "recoverGroup", groupID)

	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]

	if result.Code == "0x0" {
		return msg, nil
	}

	return msg, errors.New(GroupCodeMessage[result.Code])
}

// QueryGroupStatus 根据群组ID查询相应群组的状态
// INEXISTENT: 群组不存在
// STOPPING: 群组正在停止
// RUNNING: 群组正在运行
// STOPPED: 群组已停止
// DELETED: 群组已删除
func (api *APIHandler) QueryGroupStatus(ctx context.Context, groupID uint) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "queryGroupStatus", groupID)
	if err != nil {
		return "", err
	}

	if result.Code == "0x0" {
		return result.Status, nil
	}

	return result.Status, errors.New(GroupCodeMessage[result.Code])
}

// AddPeers 新增P2P连接配置，增加节点config.ini:[p2p]的连接配置
func (api *APIHandler) AddPeers(ctx context.Context, peers []string) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "addPeers", peers)
	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]

	if result.Code == "0x0" {
		return msg, nil
	}
	return msg, errors.New(GroupCodeMessage[result.Code])
}

// ErasePeers 删除P2P连接配置，删除节点config.ini:[p2p]的连接配置
func (api *APIHandler) ErasePeers(ctx context.Context, peers []string) (string, error) {
	var result GroupResp
	err := api.CallContext(ctx, &result, "erasePeers", peers)
	if err != nil {
		return "内部错误", err
	}

	msg := GroupCodeMessage[result.Code]

	if result.Code == "0x0" {
		return msg, nil
	}
	return msg, errors.New(GroupCodeMessage[result.Code])
}
