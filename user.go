package psn

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const USERS_URL string = "https://us-prof.np.community.playstation.net/userProfile/v1/users/"

func (oauth *oauth_response) Me() {
	client := &http.Client{}
	url := USERS_URL + "me/profile2?fields=npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount&avatarSizes=m,xl&profilePictureSizes=m,xl&languagesUsedLanguageSet=set3&psVitaTitleIcon=circled&titleIconSize=s"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	res, _ := client.Do(req)

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}
