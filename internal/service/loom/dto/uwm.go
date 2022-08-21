package dto

import (
	"time"
)

type UserWorkspaceMemberships struct {
	Data struct {
		UserWorkspaceMemberships []struct {
			CreatedAt    time.Time `json:"createdAt"`
			MemberRole   string    `json:"member_role"`
			Organization struct {
				Counts struct {
					Folders struct {
						TotalArchivedPersonalFolders int `json:"total_archived_personal_folders"`
						TotalArchivedTeamFolders     int `json:"total_archived_team_folders"`
						TotalPersonalFolders         int `json:"total_personal_folders"`
						TotalSharedFolders           int `json:"total_shared_folders"`
						TotalTeamFolders             int `json:"total_team_folders"`
					} `json:"folders"`
					Screenshots struct {
						TotalPersonalActiveScreenshots   int `json:"total_personal_active_screenshots"`
						TotalPersonalArchivedScreenshots int `json:"total_personal_archived_screenshots"`
						TotalSharedActiveScreenshots     int `json:"total_shared_active_screenshots"`
						TotalTeamActiveScreenshots       int `json:"total_team_active_screenshots"`
						TotalTeamArchivedScreenshots     int `json:"total_team_archived_screenshots"`
						TotalWorkspaceScreenshots        int `json:"total_workspace_screenshots"`
					} `json:"screenshots"`
					Spaces struct {
						TotalActiveSpaces int `json:"total_active_spaces"`
					} `json:"spaces"`
					Users struct {
						TotalActiveAdmins         int `json:"total_active_admins"`
						TotalActiveCreatorLites   int `json:"total_active_creator_lites"`
						TotalActiveCreators       int `json:"total_active_creators"`
						TotalActiveViewers        int `json:"total_active_viewers"`
						TotalActiveWorkspaceUsers int `json:"total_active_workspace_users"`
						TotalBillableMembers      int `json:"total_billable_members"`
						TotalPendingDowngrades    int `json:"total_pending_downgrades"`
						TotalWorkspaceUsers       int `json:"total_workspace_users"`
					} `json:"users"`
					Videos struct {
						SharedUnwatchedVideoCount   int `json:"shared_unwatched_video_count"`
						TotalMemberCreatedVideos    int `json:"total_member_created_videos"`
						TotalPersonalActiveVideos   int `json:"total_personal_active_videos"`
						TotalPersonalArchivedVideos int `json:"total_personal_archived_videos"`
						TotalPublishedVideos        int `json:"total_published_videos"`
						TotalSharedActiveVideos     int `json:"total_shared_active_videos"`
						TotalTeamActiveVideos       int `json:"total_team_active_videos"`
						TotalTeamArchivedVideos     int `json:"total_team_archived_videos"`
						TotalWorkspaceVideos        int `json:"total_workspace_videos"`
					} `json:"videos"`
				} `json:"counts"`
				ID     string `json:"id"`
				Limits struct {
					CreatorLiteMemberLimit int `json:"CREATOR_LITE_MEMBER_LIMIT"`
					SpaceLimit             int `json:"SPACE_LIMIT"`
				} `json:"limits"`
				MemberRoles []struct {
					Inviteable bool   `json:"inviteable"`
					IsFree     bool   `json:"is_free"`
					Name       string `json:"name"`
					Value      string `json:"value"`
				} `json:"member_roles"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"organization"`
		} `json:"userWorkspaceMemberships"`
	} `json:"data"`
	Errors []struct {
		Extensions struct {
			Code string `json:"code"`
		} `json:"extensions"`
		Locations []struct {
			Column int `json:"column"`
			Line   int `json:"line"`
		} `json:"locations"`
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors,omitempty"`
}
