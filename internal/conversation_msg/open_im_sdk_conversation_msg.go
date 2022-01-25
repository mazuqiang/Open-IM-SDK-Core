package conversation_msg

import (
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	"image"
	"open_im_sdk/open_im_sdk_callback"
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/db"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"os"
	"runtime"
	"sort"
	"sync"
)

//
//import "C"
import (
	//	//"bytes"
	//	//"encoding/gob"
	//	"encoding/json"
	//	"errors"
	//	"github.com/golang/protobuf/proto"
	//	"github.com/gorilla/websocket"
	imgtype "github.com/shamsher31/goimgtype"
	//	"image"
	//	"net/http"
	//	"open_im_sdk/pkg/db"
	//
	//	"open_im_sdk/pkg/common"
	//	"open_im_sdk/pkg/constant"
	//	"open_im_sdk/pkg/log"
	//	"open_im_sdk/pkg/sdk_params_callback"
	//	"open_im_sdk/pkg/server_api_params"
	//	"open_im_sdk/pkg/utils"
	//	"os"
	//	"sort"
	//	"sync"
	//	"time"
)

func (c *Conversation) GetAllConversationList(callback open_im_sdk_callback.Base, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "GetAllConversationList args: ")
		result := c.getAllConversationList(callback, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, "GetAllConversationList callback: ", utils.StructToJsonStringDefault(result))
	}()
}
func (c *Conversation) GetConversationListSplit(callback open_im_sdk_callback.Base, offset, count int, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "GetConversationListSplit args: ", offset, count)
		result := c.getConversationListSplit(callback, offset, count, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, "GetConversationListSplit callback: ", utils.StructToJsonStringDefault(result))
	}()
}
func (c *Conversation) SetConversationRecvMessageOpt(callback open_im_sdk_callback.Base, conversationIDList string, opt int, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "SetConversationRecvMessageOpt args: ", conversationIDList, opt)
		var unmarshalParams sdk_params_callback.SetConversationRecvMessageOptParams
		common.JsonUnmarshalCallback(conversationIDList, &unmarshalParams, callback, operationID)
		c.setConversationRecvMessageOpt(callback, unmarshalParams, opt, operationID)
		callback.OnSuccess(sdk_params_callback.SetConversationRecvMessageOptCallback)
		log.NewInfo(operationID, "SetConversationRecvMessageOpt callback: ", sdk_params_callback.AddFriendCallback)
	}()

}
func (c *Conversation) GetConversationRecvMessageOpt(callback open_im_sdk_callback.Base, conversationIDList, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "GetConversationRecvMessageOpt args: ", conversationIDList)
		var unmarshalParams sdk_params_callback.GetConversationRecvMessageOptParams
		common.JsonUnmarshalCallback(conversationIDList, &unmarshalParams, callback, operationID)
		result := c.getConversationRecvMessageOpt(callback, unmarshalParams, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, "GetConversationRecvMessageOpt callback: ", utils.StructToJsonStringDefault(result))
	}()
}
func (c *Conversation) GetOneConversation(callback open_im_sdk_callback.Base, sessionType int32, sourceID, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "GetOneConversation args: ", sessionType, sourceID)
		result := c.getOneConversation(callback, sourceID, sessionType, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, "GetRecvFriendApplicationList callback: ", utils.StructToJsonString(result))
	}()
}
func (c *Conversation) GetMultipleConversation(callback open_im_sdk_callback.Base, conversationIDList string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "GetMultipleConversation args: ", conversationIDList)
		var unmarshalParams sdk_params_callback.GetMultipleConversationParams
		common.JsonUnmarshalCallback(conversationIDList, &unmarshalParams, callback, operationID)
		result := c.getMultipleConversation(callback, unmarshalParams, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, "GetMultipleConversation callback: ", utils.StructToJsonStringDefault(result))
	}()
}
func (c *Conversation) DeleteConversation(callback open_im_sdk_callback.Base, conversationID string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "DeleteConversation args: ", conversationID)
		c.deleteConversation(callback, conversationID, operationID)
		callback.OnSuccess(sdk_params_callback.DeleteConversationCallback)
		//_ = u.triggerCmdUpdateConversation(common.updateConNode{ConID: conversationID, Action: constant.TotalUnreadMessageChanged})
		log.NewInfo(operationID, "DeleteConversation callback: ", sdk_params_callback.DeleteConversationCallback)
	}()
}
func (c *Conversation) SetConversationDraft(callback open_im_sdk_callback.Base, conversationID, draftText string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "SetConversationDraft args: ", conversationID)
		c.setConversationDraft(callback, conversationID, draftText, operationID)
		callback.OnSuccess(sdk_params_callback.SetConversationDraftCallback)
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{"", constant.NewConChange, []string{conversationID}}})
		log.NewInfo(operationID, "SetConversationDraft callback: ", sdk_params_callback.SetConversationDraftCallback)
	}()
}
func (c *Conversation) PinConversation(callback open_im_sdk_callback.Base, conversationID string, isPinned bool, operationID string) {

	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "PinConversation args: ", conversationID)
		c.pinConversation(callback, conversationID, isPinned, operationID)
		callback.OnSuccess(sdk_params_callback.PinConversationDraftCallback)
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{"", constant.NewConChange, []string{conversationID}}})
		log.NewInfo(operationID, "PinConversation callback: ", sdk_params_callback.PinConversationDraftCallback)
	}()
}
func (c *Conversation) GetTotalUnreadMsgCount(callback open_im_sdk_callback.Base, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "GetTotalUnreadMsgCount args: ")
		count, err := c.db.GetTotalUnreadMsgCount()
		common.CheckDBErrCallback(callback, err, operationID)
		callback.OnSuccess(utils.Int32ToString(count))
		log.NewInfo(operationID, "GetTotalUnreadMsgCount callback: ", utils.Int32ToString(count))
	}()
}

//
func (c *Conversation) SetConversationListener(listener open_im_sdk_callback.OnConversationListener) {
	if c.ConversationListener != nil {
		log.Error("internal", "just only set on listener")
		return
	}
	c.ConversationListener = listener
}

//

////
////func (c *Conversation) ForceSyncMsg() bool {
////	if c.syncSeq2Msg() == nil {
////		return true
////	} else {
////		return false
////	}
////}
////
////func (c *Conversation) ForceSyncJoinedGroup() {
////	u.syncJoinedGroupInfo()
////}
////
////func (c *Conversation) ForceSyncJoinedGroupMember() {
////
////	u.syncJoinedGroupMember()
////}
////
////func (c *Conversation) ForceSyncGroupRequest() {
////	u.syncGroupRequest()
////}
////
////func (c *Conversation) ForceSyncSelfGroupRequest() {
////	u.syncSelfGroupRequest()
////}
//

func (c *Conversation) CreateTextMessage(text, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Text, operationID)
	s.Content = text
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateTextAtMessage(text, atUserList, operationID string) string {
	var users []string
	_ = json.Unmarshal([]byte(atUserList), &users)
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.AtText, operationID)
	s.AtElem.Text = text
	s.AtElem.AtUserList = users
	s.Content = utils.StructToJsonString(s.AtElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateLocationMessage(description string, longitude, latitude float64, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Location, operationID)
	s.LocationElem.Description = description
	s.LocationElem.Longitude = longitude
	s.LocationElem.Latitude = latitude
	s.Content = utils.StructToJsonString(s.LocationElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateCustomMessage(data, extension string, description, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Custom, operationID)
	s.CustomElem.Data = data
	s.CustomElem.Extension = extension
	s.CustomElem.Description = description
	s.Content = utils.StructToJsonString(s.CustomElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateQuoteMessage(text string, message, operationID string) string {
	s, qs := sdk_struct.MsgStruct{}, sdk_struct.MsgStruct{}
	_ = json.Unmarshal([]byte(message), &qs)
	c.initBasicInfo(&s, constant.UserMsgType, constant.Quote, operationID)
	//Avoid nested references
	if qs.ContentType == constant.Quote {
		qs.Content = qs.QuoteElem.Text
		qs.ContentType = constant.Text
	}
	s.QuoteElem.Text = text
	s.QuoteElem.QuoteMessage = &qs
	s.Content = utils.StructToJsonString(s.QuoteElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateCardMessage(cardInfo, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Card, operationID)
	s.Content = cardInfo
	return utils.StructToJsonString(s)

}
func (c *Conversation) CreateVideoMessageFromFullPath(videoFullPath string, videoType string, duration int64, snapshotFullPath, operationID string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		dstFile := utils.FileTmpPath(videoFullPath, c.DataDir) //a->b
		s, err := utils.CopyFile(videoFullPath, dstFile)
		if err != nil {
			log.Error("internal", "open file failed: ", err, videoFullPath)
		}
		log.Info("internal", "videoFullPath dstFile", videoFullPath, dstFile, s)
		dstFile = utils.FileTmpPath(snapshotFullPath, c.DataDir) //a->b
		s, err = utils.CopyFile(snapshotFullPath, dstFile)
		if err != nil {
			log.Error("internal", "open file failed: ", err, snapshotFullPath)
		}
		log.Info("internal", "snapshotFullPath dstFile", snapshotFullPath, dstFile, s)
		wg.Done()
	}()

	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Video, operationID)
	s.VideoElem.VideoPath = videoFullPath
	s.VideoElem.VideoType = videoType
	s.VideoElem.Duration = duration
	if snapshotFullPath == "" {
		s.VideoElem.SnapshotPath = ""
	} else {
		s.VideoElem.SnapshotPath = snapshotFullPath
	}
	fi, err := os.Stat(s.VideoElem.VideoPath)
	if err != nil {
		log.Error("internal", "get file Attributes error", err.Error())
		return ""
	}
	s.VideoElem.VideoSize = fi.Size()
	if snapshotFullPath != "" {
		imageInfo, err := getImageInfo(s.VideoElem.SnapshotPath)
		if err != nil {
			log.Error("internal", "get Image Attributes error", err.Error())
			return ""
		}
		s.VideoElem.SnapshotHeight = imageInfo.Height
		s.VideoElem.SnapshotWidth = imageInfo.Width
		s.VideoElem.SnapshotSize = imageInfo.Size
	}
	wg.Wait()
	s.Content = utils.StructToJsonString(s.VideoElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateFileMessageFromFullPath(fileFullPath string, fileName, operationID string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		dstFile := utils.FileTmpPath(fileFullPath, c.DataDir)
		_, err := utils.CopyFile(fileFullPath, dstFile)
		log.Info("internal", "copy file, ", fileFullPath, dstFile)
		if err != nil {
			log.Error("internal", "open file failed: ", err.Error(), fileFullPath)
		}
		wg.Done()
	}()
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.File, operationID)
	s.FileElem.FilePath = fileFullPath
	fi, err := os.Stat(fileFullPath)
	if err != nil {
		log.Error("internal", "get file Attributes error", err.Error())
		return ""
	}
	s.FileElem.FileSize = fi.Size()
	s.FileElem.FileName = fileName
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateImageMessageFromFullPath(imageFullPath, operationID string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		dstFile := utils.FileTmpPath(imageFullPath, c.DataDir) //a->b
		_, err := utils.CopyFile(imageFullPath, dstFile)
		log.Info("internal", "copy file, ", imageFullPath, dstFile)
		if err != nil {
			log.Error("internal", "open file failed: ", err, imageFullPath)
		}
		wg.Done()
	}()

	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Picture, operationID)
	s.PictureElem.SourcePath = imageFullPath
	log.Info("internal", "ImageMessage  path:", s.PictureElem.SourcePath)
	imageInfo, err := getImageInfo(s.PictureElem.SourcePath)
	if err != nil {
		log.Error("internal", "getImageInfo err:", err.Error())
		return ""
	}
	s.PictureElem.SourcePicture.Width = imageInfo.Width
	s.PictureElem.SourcePicture.Height = imageInfo.Height
	s.PictureElem.SourcePicture.Type = imageInfo.Type
	s.PictureElem.SourcePicture.Size = imageInfo.Size
	wg.Wait()
	s.Content = utils.StructToJsonString(s.PictureElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateSoundMessageFromFullPath(soundPath string, duration int64, operationID string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		dstFile := utils.FileTmpPath(soundPath, c.DataDir) //a->b
		_, err := utils.CopyFile(soundPath, dstFile)
		log.Info("internal", "copy file, ", soundPath, dstFile)
		if err != nil {
			log.Error("internal", "open file failed: ", err, soundPath)
		}
		wg.Done()
	}()
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Voice, operationID)
	s.SoundElem.SoundPath = soundPath
	s.SoundElem.Duration = duration
	fi, err := os.Stat(s.SoundElem.SoundPath)
	if err != nil {
		log.Error("internal", "getSoundInfo err:", err.Error(), s.SoundElem.SoundPath)
		return ""
	}
	s.SoundElem.DataSize = fi.Size()
	wg.Wait()
	s.Content = utils.StructToJsonString(s.SoundElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateImageMessage(imagePath, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Picture, operationID)
	s.PictureElem.SourcePath = c.DataDir + imagePath
	log.Debug("internal", "ImageMessage  path:", s.PictureElem.SourcePath)
	imageInfo, err := getImageInfo(s.PictureElem.SourcePath)
	if err != nil {
		log.Error("internal", "get imageInfo err", err.Error())
		return ""
	}
	s.PictureElem.SourcePicture.Width = imageInfo.Width
	s.PictureElem.SourcePicture.Height = imageInfo.Height
	s.PictureElem.SourcePicture.Type = imageInfo.Type
	s.PictureElem.SourcePicture.Size = imageInfo.Size
	s.Content = utils.StructToJsonString(s.PictureElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateImageMessageByURL(sourcePicture, bigPicture, snapshotPicture, operationID string) string {
	s := sdk_struct.MsgStruct{}
	var p sdk_struct.PictureBaseInfo
	_ = json.Unmarshal([]byte(sourcePicture), &p)
	s.PictureElem.SourcePicture = p
	_ = json.Unmarshal([]byte(bigPicture), &p)
	s.PictureElem.BigPicture = p
	_ = json.Unmarshal([]byte(snapshotPicture), &p)
	s.PictureElem.SnapshotPicture = p
	c.initBasicInfo(&s, constant.UserMsgType, constant.Picture, operationID)
	s.Content = utils.StructToJsonString(s.PictureElem)
	return utils.StructToJsonString(s)
}

func msgStructToLocalChatLog(dst *db.LocalChatLog, src *sdk_struct.MsgStruct) {
	copier.Copy(dst, src)
}
func localChatLogToMsgStruct(dst *sdk_struct.NewMsgList, src []*db.LocalChatLog) {
	copier.Copy(dst, &src)
}
func (c *Conversation) checkErrAndUpdateMessage(callback open_im_sdk_callback.SendMsgCallBack, errCode int32, err error, s *sdk_struct.MsgStruct, lc *db.LocalConversation, operationID string) {
	if err != nil {
		if callback != nil {
			c.updateMsgStatusAndTriggerConversation(s.ClientMsgID, "", s.CreateTime, constant.MsgStatusSendFailed, s, lc, operationID)
			errInfo := "operationID[" + operationID + "], " + "info[" + err.Error() + "]"
			log.NewError(operationID, "checkErr ", errInfo)
			callback.OnError(errCode, errInfo)
			runtime.Goexit()
		}
	}
}
func (c *Conversation) updateMsgStatusAndTriggerConversation(clientMsgID, serverMsgID string, sendTime int64, status int32, s *sdk_struct.MsgStruct, lc *db.LocalConversation, operationID string) {
	err := c.db.UpdateMessageTimeAndStatus(clientMsgID, serverMsgID, sendTime, status)
	if err != nil {
		log.Error(operationID, "send message update message status error", clientMsgID, serverMsgID)
	}
	s.SendTime = sendTime
	s.Status = status
	s.ServerMsgID = serverMsgID
	lc.LatestMsg = utils.StructToJsonString(s)
	lc.LatestMsgSendTime = sendTime
	log.Info(operationID, "2 send message come here", *lc)
	//会话数据库操作，触发UI会话回调
	_ = common.TriggerCmdUpdateConversation(common.UpdateConNode{ConID: lc.ConversationID, Action: constant.AddConOrUpLatMsg, Args: lc}, c.ch)

}
func (c *Conversation) getUserNameAndFaceUrlByUid(callback open_im_sdk_callback.SendMsgCallBack, friendUserID, operationID string) (faceUrl, name string, err error) {
	friendInfo, err := c.db.GetFriendInfoByFriendUserID(friendUserID)
	if err == nil {
		if friendInfo.Remark != "" {
			return friendInfo.FaceURL, friendInfo.Remark, nil
		} else {
			return friendInfo.FaceURL, friendInfo.Nickname, nil
		}
	} else {
		userInfos := c.user.GetUsersInfoFromSvr(callback, []string{friendUserID}, operationID)
		for _, v := range userInfos {
			return v.FaceURL, v.Nickname, nil
		}
	}
	return "", "", errors.New("getUserNameAndFaceUrlByUid err")

}

func (c *Conversation) SendMessage(callback open_im_sdk_callback.SendMsgCallBack, message, recvID, groupID string, offlinePushInfo string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		//参数校验
		s := sdk_struct.MsgStruct{}
		common.JsonUnmarshalAndArgsValidate(message, &s, callback, operationID)
		p := server_api_params.OfflinePushInfo{}
		common.JsonUnmarshalAndArgsValidate(offlinePushInfo, &p, callback, operationID)
		if recvID == "" && groupID == "" {
			common.CheckAnyErrCallback(callback, 201, errors.New("recvID && groupID not be allowed"), operationID)
		}
		var localMessage db.LocalChatLog
		var conversationID string
		options := make(map[string]bool, 2)
		lc := &db.LocalConversation{LatestMsgSendTime: s.CreateTime}
		//根据单聊群聊类型组装消息和会话
		if recvID == "" {
			s.SessionType = constant.GroupChatType
			s.GroupID = groupID
			conversationID = c.GetConversationIDBySessionType(groupID, constant.GroupChatType)
			lc.GroupID = groupID
			lc.ConversationType = constant.GroupChatType
			g, err := c.db.GetGroupInfoByGroupID(groupID)
			common.CheckAnyErrCallback(callback, 202, err, operationID)
			lc.ShowName = g.GroupName
			lc.FaceURL = g.FaceURL
			groupMemberUidList, err := c.db.GetGroupMemberUIDListByGroupID(groupID)
			common.CheckAnyErrCallback(callback, 202, err, operationID)
			if !utils.IsContain(s.SendID, groupMemberUidList) {
				common.CheckAnyErrCallback(callback, 208, errors.New("you not exist in this group"), operationID)
			}
		} else {
			s.SessionType = constant.SingleChatType
			s.RecvID = recvID
			conversationID = c.GetConversationIDBySessionType(recvID, constant.SingleChatType)
			lc.UserID = recvID
			lc.ConversationType = constant.SingleChatType
			faceUrl, name, err := c.getUserNameAndFaceUrlByUid(callback, recvID, operationID)
			common.CheckAnyErrCallback(callback, 301, err, operationID)
			lc.FaceURL = faceUrl
			lc.ShowName = name
		}
		lc.ConversationID = conversationID
		lc.LatestMsg = utils.StructToJsonString(s)
		msgStructToLocalChatLog(&localMessage, &s)
		err := c.db.InsertMessage(&localMessage)
		common.CheckAnyErrCallback(callback, 201, err, operationID)
		log.Info(operationID, "send message come here", *lc)
		_ = common.TriggerCmdUpdateConversation(common.UpdateConNode{ConID: conversationID, Action: constant.AddConOrUpLatMsg, Args: lc}, c.ch)
		var delFile []string
		//media file handle
		switch s.ContentType {
		case constant.Picture:
			var sourcePath string
			if utils.FileExist(s.PictureElem.SourcePath) {
				sourcePath = s.PictureElem.SourcePath
				delFile = append(delFile, utils.FileTmpPath(s.PictureElem.SourcePath, c.DataDir))
			} else {
				sourcePath = utils.FileTmpPath(s.PictureElem.SourcePath, c.DataDir)
				delFile = append(delFile, sourcePath)
			}
			log.Info(operationID, "file", sourcePath, delFile)
			sourceUrl, uuid, err := c.UploadImage(sourcePath, callback.OnProgress)
			c.checkErrAndUpdateMessage(callback, 301, err, &s, lc, operationID)
			s.PictureElem.SourcePicture.Url = sourceUrl
			s.PictureElem.SourcePicture.UUID = uuid
			s.PictureElem.SnapshotPicture.Url = sourceUrl + "?imageView2/2/w/" + constant.ZoomScale + "/h/" + constant.ZoomScale
			s.PictureElem.SnapshotPicture.Width = int32(utils.StringToInt(constant.ZoomScale))
			s.PictureElem.SnapshotPicture.Height = int32(utils.StringToInt(constant.ZoomScale))
			s.Content = utils.StructToJsonString(s.PictureElem)

		case constant.Voice:
			var sourcePath string
			if utils.FileExist(s.SoundElem.SoundPath) {
				sourcePath = s.SoundElem.SoundPath
				delFile = append(delFile, utils.FileTmpPath(s.SoundElem.SoundPath, c.DataDir))
			} else {
				sourcePath = utils.FileTmpPath(s.SoundElem.SoundPath, c.DataDir)
				delFile = append(delFile, sourcePath)
			}
			log.Info(operationID, "file", sourcePath, delFile)
			soundURL, uuid, err := c.UploadSound(sourcePath, callback.OnProgress)
			c.checkErrAndUpdateMessage(callback, 301, err, &s, lc, operationID)
			s.SoundElem.SourceURL = soundURL
			s.SoundElem.UUID = uuid
			s.Content = utils.StructToJsonString(s.SoundElem)

		case constant.Video:
			var videoPath string
			var snapPath string
			if utils.FileExist(s.VideoElem.VideoPath) {
				videoPath = s.VideoElem.VideoPath
				snapPath = s.VideoElem.SnapshotPath
				delFile = append(delFile, utils.FileTmpPath(s.VideoElem.VideoPath, c.DataDir))
				delFile = append(delFile, utils.FileTmpPath(s.VideoElem.SnapshotPath, c.DataDir))
			} else {
				videoPath = utils.FileTmpPath(s.VideoElem.VideoPath, c.DataDir)
				snapPath = utils.FileTmpPath(s.VideoElem.SnapshotPath, c.DataDir)
				delFile = append(delFile, videoPath)
				delFile = append(delFile, snapPath)
			}
			log.Info(operationID, "file: ", videoPath, snapPath, delFile)
			snapshotURL, snapshotUUID, videoURL, videoUUID, err := c.UploadVideo(videoPath, snapPath, callback.OnProgress)
			c.checkErrAndUpdateMessage(callback, 301, err, &s, lc, operationID)
			s.VideoElem.VideoURL = videoURL
			s.VideoElem.SnapshotUUID = snapshotUUID
			s.VideoElem.SnapshotURL = snapshotURL
			s.VideoElem.VideoUUID = videoUUID
			s.Content = utils.StructToJsonString(s.VideoElem)
		case constant.File:
			fileURL, fileUUID, err := c.UploadFile(s.FileElem.FilePath, callback.OnProgress)
			c.checkErrAndUpdateMessage(callback, 301, err, &s, lc, operationID)
			s.FileElem.SourceURL = fileURL
			s.FileElem.UUID = fileUUID
			s.Content = utils.StructToJsonString(s.FileElem)
		case constant.Text:
		case constant.AtText:
		case constant.Location:
		case constant.Custom:
		case constant.Merger:
		case constant.Quote:
		case constant.Card:
		default:
			common.CheckAnyErrCallback(callback, 202, errors.New("contentType not currently supported"+utils.Int32ToString(s.ContentType)), operationID)
		}
		msgStructToLocalChatLog(&localMessage, &s)
		err = c.db.UpdateMessage(&localMessage)
		common.CheckAnyErrCallback(callback, 201, err, operationID)
		c.sendMessageToServer(&s, lc, callback, delFile, &p, options, operationID)
	}()
}
func (c *Conversation) SendMessageNotOss(callback open_im_sdk_callback.SendMsgCallBack, message, recvID, groupID string, offlinePushInfo string, operationID string) {
	go func() {
		s := sdk_struct.MsgStruct{}
		common.JsonUnmarshalAndArgsValidate(message, &s, callback, operationID)
		p := server_api_params.OfflinePushInfo{}
		common.JsonUnmarshalAndArgsValidate(offlinePushInfo, &p, callback, operationID)
		if recvID == "" && groupID == "" {
			common.CheckAnyErrCallback(callback, 201, errors.New("recvID && groupID not be allowed"), operationID)
		}
		var localMessage db.LocalChatLog
		var conversationID string
		var options map[string]bool
		lc := db.LocalConversation{
			LatestMsgSendTime: s.CreateTime,
		}
		//根据单聊群聊类型组装消息和会话
		if recvID == "" {
			s.SessionType = constant.GroupChatType
			s.GroupID = groupID
			conversationID = c.GetConversationIDBySessionType(groupID, constant.GroupChatType)
			lc.GroupID = groupID
			lc.ConversationType = constant.GroupChatType
			g, err := c.db.GetGroupInfoByGroupID(groupID)
			common.CheckAnyErrCallback(callback, 202, err, operationID)
			lc.ShowName = g.GroupName
			lc.FaceURL = g.FaceURL
			groupMemberUidList, err := c.db.GetGroupMemberUIDListByGroupID(groupID)
			common.CheckAnyErrCallback(callback, 202, err, operationID)
			if !utils.IsContain(s.SendID, groupMemberUidList) {
				common.CheckAnyErrCallback(callback, 208, errors.New("you not exist in this group"), operationID)
			}
		} else {
			s.SessionType = constant.SingleChatType
			s.RecvID = recvID
			conversationID = c.GetConversationIDBySessionType(recvID, constant.SingleChatType)
			lc.UserID = recvID
			lc.ConversationType = constant.SingleChatType
			faceUrl, name, err := c.getUserNameAndFaceUrlByUid(callback, recvID, operationID)
			common.CheckAnyErrCallback(callback, 301, err, operationID)
			lc.FaceURL = faceUrl
			lc.ShowName = name
		}
		lc.ConversationID = conversationID
		lc.LatestMsg = utils.StructToJsonString(s)
		msgStructToLocalChatLog(&localMessage, &s)
		err := c.db.InsertMessage(&localMessage)
		common.CheckAnyErrCallback(callback, 201, err, operationID)
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{conversationID, constant.AddConOrUpLatMsg,
		//c}})
		//u.doUpdateConversation(cmd2Value{Value: updateConNode{"", NewConChange, []string{conversationID}}})
		//_ = u.triggerCmdUpdateConversation(updateConNode{conversationID, ConChange, ""})
		options = make(map[string]bool, 2)
		var delFile []string

		msgStructToLocalChatLog(&localMessage, &s)
		err = c.db.UpdateMessage(&localMessage)
		common.CheckAnyErrCallback(callback, 201, err, operationID)
		c.sendMessageToServer(&s, &lc, callback, delFile, &p, options, operationID)

	}()
}
func (c *Conversation) internalSendMessage(callback open_im_sdk_callback.Base, s *sdk_struct.MsgStruct, recvID, groupID, operationID string, p *server_api_params.OfflinePushInfo, onlineUserOnly bool, options map[string]bool) error {
	if recvID == "" && groupID == "" {
		common.CheckAnyErrCallback(callback, 201, errors.New("recvID && groupID not be allowed"), operationID)
	}
	if recvID == "" {
		s.SessionType = constant.GroupChatType
		s.GroupID = groupID
		groupMemberUidList, err := c.db.GetGroupMemberUIDListByGroupID(groupID)
		common.CheckAnyErrCallback(callback, 202, err, operationID)
		if !utils.IsContain(s.SendID, groupMemberUidList) {
			common.CheckAnyErrCallback(callback, 208, errors.New("you not exist in this group"), operationID)
		}

	} else {
		s.SessionType = constant.SingleChatType
		s.RecvID = recvID
	}

	if onlineUserOnly {
		options[constant.IsHistory] = false
		options[constant.IsPersistent] = false
		options[constant.IsOfflinePush] = false
	}

	var wsMsgData server_api_params.MsgData
	copier.Copy(&wsMsgData, s)
	wsMsgData.Content = []byte(s.Content)
	wsMsgData.CreateTime = int64(s.CreateTime)
	wsMsgData.Options = options
	wsMsgData.OfflinePushInfo = p
	timeout := 300
	retryTimes := 0
	_, err := c.SendReqWaitResp(&wsMsgData, constant.WSSendMsg, timeout, retryTimes, c.loginUserID, operationID)
	common.CheckAnyErrCallback(callback, 301, err, operationID)
	return nil

}
func (c *Conversation) sendMessageToServer(s *sdk_struct.MsgStruct, lc *db.LocalConversation, callback open_im_sdk_callback.SendMsgCallBack,
	delFile []string, offlinePushInfo *server_api_params.OfflinePushInfo, options map[string]bool, operationID string) {
	//Protocol conversion
	var wsMsgData server_api_params.MsgData
	copier.Copy(&wsMsgData, s)
	wsMsgData.Content = []byte(s.Content)
	wsMsgData.CreateTime = s.CreateTime
	wsMsgData.Options = options
	wsMsgData.OfflinePushInfo = offlinePushInfo
	timeout := 300
	retryTimes := 60
	resp, err := c.SendReqWaitResp(&wsMsgData, constant.WSSendMsg, timeout, retryTimes, c.loginUserID, operationID)
	c.checkErrAndUpdateMessage(callback, 302, err, s, lc, operationID)
	callback.OnProgress(100)
	callback.OnSuccess("")
	//remove media cache file
	for _, v := range delFile {
		err := os.Remove(v)
		if err != nil {
			log.Error(operationID, "remove failed,", err.Error(), v)
		}
		log.Debug(operationID, "remove file: ", v)
	}
	var sendMsgResp server_api_params.UserSendMsgResp
	_ = proto.Unmarshal(resp.Data, &sendMsgResp)
	c.updateMsgStatusAndTriggerConversation(sendMsgResp.ClientMsgID, sendMsgResp.ServerMsgID, sendMsgResp.SendTime, constant.MsgStatusSendSuccess, s, lc, operationID)

}

func (c *Conversation) CreateSoundMessageByURL(soundBaseInfo, operationID string) string {
	s := sdk_struct.MsgStruct{}
	var soundElem sdk_struct.SoundBaseInfo
	_ = json.Unmarshal([]byte(soundBaseInfo), &soundElem)
	s.SoundElem = soundElem
	c.initBasicInfo(&s, constant.UserMsgType, constant.Voice, operationID)
	s.Content = utils.StructToJsonString(s.SoundElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateSoundMessage(soundPath string, duration int64, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Voice, operationID)
	s.SoundElem.SoundPath = c.DataDir + soundPath
	s.SoundElem.Duration = duration
	fi, err := os.Stat(s.SoundElem.SoundPath)
	if err != nil {
		log.Error("internal", "get sound info err", err.Error())
		return ""
	}
	s.SoundElem.DataSize = fi.Size()
	s.Content = utils.StructToJsonString(s.SoundElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateVideoMessageByURL(videoBaseInfo, operationID string) string {
	s := sdk_struct.MsgStruct{}
	var videoElem sdk_struct.VideoBaseInfo
	_ = json.Unmarshal([]byte(videoBaseInfo), &videoElem)
	s.VideoElem = videoElem
	c.initBasicInfo(&s, constant.UserMsgType, constant.Video, operationID)
	s.Content = utils.StructToJsonString(s.VideoElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateVideoMessage(videoPath string, videoType string, duration int64, snapshotPath, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.Video, operationID)
	s.VideoElem.VideoPath = c.DataDir + videoPath
	s.VideoElem.VideoType = videoType
	s.VideoElem.Duration = duration
	if snapshotPath == "" {
		s.VideoElem.SnapshotPath = ""
	} else {
		s.VideoElem.SnapshotPath = c.DataDir + snapshotPath
	}
	fi, err := os.Stat(s.VideoElem.VideoPath)
	if err != nil {
		log.Error("internal", "get video file error", err.Error())
		return ""
	}
	s.VideoElem.VideoSize = fi.Size()
	if snapshotPath != "" {
		imageInfo, err := getImageInfo(s.VideoElem.SnapshotPath)
		if err != nil {
			log.Error("internal", "get snapshot info ", err.Error())
			return ""
		}
		s.VideoElem.SnapshotHeight = imageInfo.Height
		s.VideoElem.SnapshotWidth = imageInfo.Width
		s.VideoElem.SnapshotSize = imageInfo.Size
	}
	s.Content = utils.StructToJsonString(s.VideoElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateFileMessageByURL(fileBaseInfo, operationID string) string {
	s := sdk_struct.MsgStruct{}
	var fileElem sdk_struct.FileBaseInfo
	_ = json.Unmarshal([]byte(fileBaseInfo), &fileElem)
	s.FileElem = fileElem
	c.initBasicInfo(&s, constant.UserMsgType, constant.File, operationID)
	s.Content = utils.StructToJsonString(s.FileElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateFileMessage(filePath string, fileName, operationID string) string {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(&s, constant.UserMsgType, constant.File, operationID)
	s.FileElem.FilePath = c.DataDir + filePath
	s.FileElem.FileName = fileName
	fi, err := os.Stat(s.FileElem.FilePath)
	if err != nil {
		log.Error("internal", "get file message err", err.Error())
		return ""
	}
	s.FileElem.FileSize = fi.Size()
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateMergerMessage(messageList, title, summaryList, operationID string) string {
	var messages []*sdk_struct.MsgStruct
	var summaries []string
	s := sdk_struct.MsgStruct{}
	err := json.Unmarshal([]byte(messageList), &messages)
	if err != nil {
		log.Error("internal", "messages Unmarshal err", err.Error())
		return ""
	}
	_ = json.Unmarshal([]byte(summaryList), &summaries)
	c.initBasicInfo(&s, constant.UserMsgType, constant.Merger, operationID)
	s.MergeElem.AbstractList = summaries
	s.MergeElem.Title = title
	s.MergeElem.MultiMessage = messages
	s.Content = utils.StructToJsonString(s.MergeElem)
	return utils.StructToJsonString(s)
}
func (c *Conversation) CreateForwardMessage(m, operationID string) string {
	s := sdk_struct.MsgStruct{}
	err := json.Unmarshal([]byte(m), &s)
	if err != nil {
		log.Error("internal", "messages Unmarshal err", err.Error())
		return ""
	}
	if s.Status != constant.MsgStatusSendSuccess {
		log.Error("internal", "only send success message can be revoked")
		return ""
	}
	c.initBasicInfo(&s, constant.UserMsgType, s.ContentType, operationID)
	//Forward message seq is set to 0
	s.Seq = 0
	return utils.StructToJsonString(s)
}
func (c *Conversation) GetHistoryMessageList(callback open_im_sdk_callback.Base, getMessageOptions, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		var messageList sdk_struct.NewMsgList
		log.NewInfo(operationID, "GetHistoryMessageList args: ", getMessageOptions)
		var unmarshalParams sdk_params_callback.GetHistoryMessageListParams
		common.JsonUnmarshalCallback(getMessageOptions, &unmarshalParams, callback, operationID)
		result := c.getHistoryMessageList(callback, unmarshalParams, operationID)
		localChatLogToMsgStruct(&messageList, result)
		for _, v := range messageList {
			err := c.msgHandleByContentType(v)
			if err != nil {
				log.Error(operationID, "Parsing data error:", err.Error(), v)
				continue
			}
		}
		sort.Sort(messageList)
		callback.OnSuccess(utils.StructToJsonStringDefault(messageList))
		log.NewInfo(operationID, "GetHistoryMessageList callback: ", utils.StructToJsonStringDefault(messageList))
	}()
}

func (c *Conversation) RevokeMessage(callback open_im_sdk_callback.Base, message string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "RevokeMessage args: ", message)
		var unmarshalParams sdk_params_callback.RevokeMessageParams
		common.JsonUnmarshalCallback(message, &unmarshalParams, callback, operationID)
		c.revokeOneMessage(callback, unmarshalParams, operationID)
		callback.OnSuccess(sdk_params_callback.RevokeMessageCallback)
		log.NewInfo(operationID, "RevokeMessage callback: ", sdk_params_callback.RevokeMessageCallback)
	}()
}
func (c *Conversation) TypingStatusUpdate(callback open_im_sdk_callback.Base, recvID, msgTip, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "TypingStatusUpdate args: ", recvID, msgTip)
		c.typingStatusUpdate(callback, recvID, msgTip, operationID)
		callback.OnSuccess(sdk_params_callback.TypingStatusUpdateCallback)
		log.NewInfo(operationID, "TypingStatusUpdate callback: ", sdk_params_callback.TypingStatusUpdateCallback)
	}()
}

func (c *Conversation) MarkC2CMessageAsRead(callback open_im_sdk_callback.Base, recvID string, msgIDList, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "MarkC2CMessageAsRead args: ", recvID, msgIDList)
		if len(msgIDList) == 0 {
			//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{ConID: conversationID, Action: constant.UnreadCountSetZero}})
			//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{"", constant.NewConChange, []string{conversationID}}})
			callback.OnSuccess(sdk_params_callback.MarkC2CMessageAsReadCallback)
			return
		}
		c.markC2CMessageAsRead(callback, msgIDList, recvID, operationID)
		callback.OnSuccess(sdk_params_callback.MarkC2CMessageAsReadCallback)
		log.NewInfo(operationID, "MarkC2CMessageAsRead callback: ", sdk_params_callback.MarkC2CMessageAsReadCallback)
	}()

}

////Deprecated
func (c *Conversation) MarkSingleMessageHasRead(callback open_im_sdk_callback.Base, userID string) {
	go func() {
		//conversationID := c.GetConversationIDBySessionType(userID, constant.SingleChatType)
		//if err := u.setSingleMessageHasRead(userID); err != nil {
		//	callback.OnError(201, err.Error())
		//} else {
		callback.OnSuccess("")
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{ConID: conversationID, Action: constant.UnreadCountSetZero}})
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{"", constant.NewConChange, []string{conversationID}}})
		//}
	}()
}

//func (c *Conversation) MarkAllConversationHasRead(callback common.Base, userID string) {
//	go func() {
//		conversationID := utils.GetConversationIDBySessionType(userID, constant.SingleChatType)
//		//if err := u.setSingleMessageHasRead(userID); err != nil {
//		//	callback.OnError(201, err.Error())
//		//} else {
//		callback.OnSuccess("")
//		u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{ConID: conversationID, Action: constant.UnreadCountSetZero}})
//		u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{"", constant.NewConChange, []string{conversationID}}})
//		//}
//	}()
//}

func (c *Conversation) MarkGroupMessageHasRead(callback open_im_sdk_callback.Base, groupID string, operationID string) {
	go func() {
		//conversationID := c.GetConversationIDBySessionType(groupID, constant.GroupChatType)
		//if err := u.setGroupMessageHasRead(groupID); err != nil {
		//	callback.OnError(201, err.Error())
		//} else {
		callback.OnSuccess("")
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{ConID: conversationID, Action: constant.UnreadCountSetZero}})
		//u.doUpdateConversation(common.cmd2Value{Value: common.updateConNode{"", constant.NewConChange, []string{conversationID}}})
		//}
	}()
}

func (c *Conversation) DeleteMessageFromLocalStorage(callback open_im_sdk_callback.Base, message string, operationID string) {
	go func() {
		s := sdk_struct.MsgStruct{}
		common.JsonUnmarshalAndArgsValidate(message, &s, callback, operationID)
		c.deleteMessageFromLocalStorage(callback, &s, operationID)
		callback.OnSuccess("")
	}()
}

func (c *Conversation) ClearC2CHistoryMessage(callback open_im_sdk_callback.Base, userID string, operationID string) {
	go func() {
		c.clearC2CHistoryMessage(callback, userID, operationID)
		callback.OnSuccess("")

	}()
}

func (c *Conversation) ClearGroupHistoryMessage(callback open_im_sdk_callback.Base, groupID string, operationID string) {
	go func() {
		c.clearGroupHistoryMessage(callback, groupID, operationID)
		callback.OnSuccess("")

	}()
}

func (c *Conversation) InsertSingleMessageToLocalStorage(callback open_im_sdk_callback.Base, message, userID, sendID, operationID string) string {
	s := sdk_struct.MsgStruct{}
	common.JsonUnmarshalAndArgsValidate(message, &s, callback, operationID)
	localMessage := db.LocalChatLog{}
	msgStructToLocalChatLog(&localMessage, &s)
	s.SendID = sendID
	s.RecvID = userID
	s.ClientMsgID = utils.GetMsgID(s.SendID)
	s.SendTime = utils.GetCurrentTimestampByMill()

	go func() {
		clientMsgID := c.insertMessageToLocalStorage(callback, &localMessage, operationID)
		callback.OnSuccess(clientMsgID)
	}()
	return s.ClientMsgID
}

func (c *Conversation) InsertGroupMessageToLocalStorage(callback open_im_sdk_callback.Base, message, groupID, sendID, operationID string) string {
	s := sdk_struct.MsgStruct{}
	common.JsonUnmarshalAndArgsValidate(message, &s, callback, operationID)
	localMessage := db.LocalChatLog{}
	msgStructToLocalChatLog(&localMessage, &s)
	s.SendID = sendID
	s.RecvID = groupID
	s.ClientMsgID = utils.GetMsgID(s.SendID)
	s.SendTime = utils.GetCurrentTimestampByMill()

	go func() {
		clientMsgID := c.insertMessageToLocalStorage(callback, &localMessage, operationID)
		callback.OnSuccess(clientMsgID)
	}()
	return s.ClientMsgID
}

//func (c *Conversation) FindMessages(callback common.Base, messageIDList string) {
//	go func() {
//		var c []string
//		err := json.Unmarshal([]byte(messageIDList), &c)
//		if err != nil {
//			callback.OnError(200, err.Error())
//			utils.sdkLog("Unmarshal failed, ", err.Error())
//
//		}
//		err, list := u.getMultipleMessageModel(c)
//		if err != nil {
//			callback.OnError(203, err.Error())
//		} else {
//			if list != nil {
//				callback.OnSuccess(utils.structToJsonString(list))
//			} else {
//				callback.OnSuccess(utils.structToJsonString([]utils.MsgStruct{}))
//			}
//		}
//	}()
//}

func getImageInfo(filePath string) (*sdk_struct.ImageInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, utils.Wrap(err, "open file err")
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, utils.Wrap(err, "image file  Decode err")
	}

	datatype, err := imgtype.Get(filePath)
	if err != nil {
		return nil, utils.Wrap(err, "image file  get type err")
	}
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, utils.Wrap(err, "image file  Stat err")
	}

	b := img.Bounds()

	return &sdk_struct.ImageInfo{int32(b.Max.X), int32(b.Max.Y), datatype, fi.Size()}, nil

}
