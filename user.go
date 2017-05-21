package psn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user_error struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type user_profile struct {
	Profile struct {
		OnlineID   string `json:"onlineId"`
		NpID       string `json:"npId"`
		AvatarUrls []struct {
			Size      string `json:"size"`
			AvatarURL string `json:"avatarUrl"`
		} `json:"avatarUrls"`
		Plus          int      `json:"plus"`
		AboutMe       string   `json:"aboutMe"`
		LanguagesUsed []string `json:"languagesUsed"`
		TrophySummary struct {
			Level          int `json:"level"`
			Progress       int `json:"progress"`
			EarnedTrophies struct {
				Platinum int `json:"platinum"`
				Gold     int `json:"gold"`
				Silver   int `json:"silver"`
				Bronze   int `json:"bronze"`
			} `json:"earnedTrophies"`
		} `json:"trophySummary"`
		IsOfficiallyVerified                    bool   `json:"isOfficiallyVerified"`
		PersonalDetailSharing                   string `json:"personalDetailSharing"`
		PersonalDetailSharingRequestMessageFlag bool   `json:"personalDetailSharingRequestMessageFlag"`
		PrimaryOnlineStatus                     string `json:"primaryOnlineStatus"`
		Presences                               []struct {
			OnlineStatus     string `json:"onlineStatus"`
			HasBroadcastData bool   `json:"hasBroadcastData"`
		} `json:"presences"`
		FriendRelation      string `json:"friendRelation"`
		RequestMessageFlag  bool   `json:"requestMessageFlag"`
		Blocking            bool   `json:"blocking"`
		FriendsCount        int    `json:"friendsCount"`
		MutualFriendsCount  int    `json:"mutualFriendsCount"`
		Following           bool   `json:"following"`
		FollowingUsersCount int    `json:"followingUsersCount"`
		FollowerCount       int    `json:"followerCount"`
	} `json:"profile"`
}

//Used for debugging API responses
//data, _ := ioutil.ReadAll(resp.Body)
//fmt.Printf("%s", string(data))

const USERS_URL string = "https://us-prof.np.community.playstation.net/userProfile/v1/users/"

func (oauth *oauth_response) Me() (user_profile, error) {
	client := &http.Client{}
	url := USERS_URL + "me/profile2?fields=npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount&avatarSizes=m,xl&profilePictureSizes=m,xl&languagesUsedLanguageSet=set3&psVitaTitleIcon=circled&titleIconSize=s"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return user_profile{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return user_profile{}, err
		}
		return user_profile{}, fmt.Errorf(api_error.Error.Message)
	}

	var profile user_profile
	err := json.NewDecoder(resp.Body).Decode(&profile)

	if err != nil {
		return user_profile{}, err
	}

	return profile, nil
}
