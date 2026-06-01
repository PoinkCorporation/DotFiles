package tests

import (
	"testing"

	chatv1 "github.com/PoinkCorporation/Chat-service/gen/go"
	"github.com/PoinkCorporation/Chat-service/tests/suite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListUserChats_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	userID := randomID()
	title1 := randomTitle()

	resp1, err := st.ChatClient.CreateChat(ctx, &chatv1.CreateChatRequest{
		SenderId: userID,
		Type:     "group",
		Title:    &title1,
	})
	require.NoError(t, err)

	resp2, err := st.ChatClient.CreateChat(ctx, &chatv1.CreateChatRequest{
		SenderId: userID,
		Type:     "cloud private",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		// _, _ = st.ChatClient.DeleteChat(ctx, &chatv1.DeleteChatRequest{Id: resp1.GetId(), SenderId: userID})
		// _, _ = st.ChatClient.DeleteChat(ctx, &chatv1.DeleteChatRequest{Id: resp2.GetId(), SenderId: userID})
	})

	listResp, err := st.ChatClient.ListUserChats(ctx, &chatv1.ListUserChatsRequest{UserId: userID})
	require.NoError(t, err)
	require.Len(t, listResp.GetChats(), 2)

	foundIDs := make(map[int64]bool)
	for _, ch := range listResp.GetChats() {
		foundIDs[ch.GetId()] = true
	}
	assert.True(t, foundIDs[resp1.GetId()], "chat 1 should be in list")
	assert.True(t, foundIDs[resp2.GetId()], "chat 2 should be in list")
}

func TestListUserChats_Empty(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.ChatClient.ListUserChats(ctx, &chatv1.ListUserChatsRequest{UserId: randomID()})
	require.NoError(t, err)
	assert.Empty(t, resp.GetChats())
}

func TestListUserChats_Fail_NoUserID(t *testing.T) {
	ctx, st := suite.New(t)

	_, err := st.ChatClient.ListUserChats(ctx, &chatv1.ListUserChatsRequest{})
	require.Error(t, err)

	stat, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, stat.Code())
}

func TestUpdateChat_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	senderID := randomID()
	originalTitle := randomTitle()
	newTitle := randomTitle()

	created, err := st.ChatClient.CreateChat(ctx, &chatv1.CreateChatRequest{
		SenderId: senderID,
		Type:     "group",
		Title:    &originalTitle,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = st.ChatClient.DeleteChat(ctx, &chatv1.DeleteChatRequest{Id: created.GetId(), SenderId: senderID})
	})

	updated, err := st.ChatClient.UpdateChat(ctx, &chatv1.UpdateChatRequest{
		Id:    created.GetId(),
		Title: &newTitle,
	})
	require.NoError(t, err)

	assert.Equal(t, newTitle, *updated.Title)
	assert.Equal(t, "group", updated.GetType())
	assert.Equal(t, created.GetCreatedBy(), updated.GetCreatedBy())
}

func TestUpdateChat_ChangeType(t *testing.T) {
	ctx, st := suite.New(t)

	senderID := randomID()
	title := randomTitle()

	created, err := st.ChatClient.CreateChat(ctx, &chatv1.CreateChatRequest{
		SenderId: senderID,
		Type:     "group",
		Title:    &title,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = st.ChatClient.DeleteChat(ctx, &chatv1.DeleteChatRequest{Id: created.GetId(), SenderId: senderID})
	})

	newType := "e2e private"
	updated, err := st.ChatClient.UpdateChat(ctx, &chatv1.UpdateChatRequest{
		Id:   created.GetId(),
		Type: &newType,
	})
	require.NoError(t, err)

	assert.Equal(t, newType, updated.GetType())
	assert.Equal(t, title, *updated.Title)
}

func TestUpdateChat_Fail_NoID(t *testing.T) {
	ctx, st := suite.New(t)

	_, err := st.ChatClient.UpdateChat(ctx, &chatv1.UpdateChatRequest{})
	require.Error(t, err)

	stat, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, stat.Code())
}

func TestDeleteChat_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	senderID := randomID()
	created, err := st.ChatClient.CreateChat(ctx, &chatv1.CreateChatRequest{
		SenderId: senderID,
		Type:     "cloud private",
	})
	require.NoError(t, err)

	_, err = st.ChatClient.DeleteChat(ctx, &chatv1.DeleteChatRequest{
		Id:       created.GetId(),
		SenderId: senderID,
	})
	require.NoError(t, err)
}

func TestDeleteChat_Fail_NoID(t *testing.T) {
	ctx, st := suite.New(t)

	_, err := st.ChatClient.DeleteChat(ctx, &chatv1.DeleteChatRequest{})
	require.Error(t, err)

	stat, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, stat.Code())
}
