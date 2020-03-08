// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSBookmark undocumented
type IOSBookmark struct {
	// Object is the base model of IOSBookmark
	Object
	// URL URL allowed to access
	URL *string `json:"url,omitempty"`
	// BookmarkFolder The folder into which the bookmark should be added in Safari
	BookmarkFolder *string `json:"bookmarkFolder,omitempty"`
	// DisplayName The display name of the bookmark
	DisplayName *string `json:"displayName,omitempty"`
}