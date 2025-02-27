package core

import (
	"regexp"
)

type route struct {
	pattern string
	method  string
	action  string
}

var routes = []route{
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#build-image-from-a-dockerfile
	{pattern: "/build", method: "POST", action: ActionImageBuild},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#create-a-new-image-from-a-container-s-changes
	{pattern: "/commit", method: "POST", action: ActionContainerCommit},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#monitor-docker-s-events
	{pattern: "/events", method: "GET", action: ActionDockerEvents},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#show-the-docker-version-information
	{pattern: "/version", method: "GET", action: ActionDockerVersion},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.20/#check-auth-configuration
	{pattern: "/auth", method: "POST", action: ActionDockerCheckAuth},
	// https://docs.docker.com/engine/api/v1.37/#operation/SecretList
	{pattern: "/secrets", method: "GET", action: ActionSecretList},
	// https://docs.docker.com/engine/api/v1.37/#operation/SecretInspect
	{pattern: "/secrets/.+", method: "GET", action: ActionSecretInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#wait-a-container
	{pattern: "/containers/.+/wait", method: "POST", action: ActionContainerWait},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#resize-a-container-tty
	{pattern: "/containers/.+/resize", method: "POST", action: ActionContainerResize},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#export-a-container
	{pattern: "/containers/.+/export", method: "POST", action: ActionContainerExport},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#export-a-container
	{pattern: "/containers/.+/stop", method: "POST", action: ActionContainerStop},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#kill-a-container
	{pattern: "/containers/.*/kill", method: "POST", action: ActionContainerKill},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#restart-a-container
	{pattern: "/containers/.+/restart", method: "POST", action: ActionContainerRestart},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#start-a-container
	{pattern: "/containers/.+/start", method: "POST", action: ActionContainerStart},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-create
	{pattern: "/containers/.+/exec", method: "POST", action: ActionContainerExecCreate},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#unpause-a-container
	{pattern: "/containers/.+/unpause", method: "POST", action: ActionContainerUnpause},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#pause-a-container
	{pattern: "/containers/.+/pause", method: "POST", action: ActionContainerPause},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#copy-files-or-folders-from-a-container
	{pattern: "/containers/.+/copy", method: "POST", action: ActionContainerCopyFiles},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#extract-an-archive-of-files-or-folders-to-a-directory-in-a-container
	{pattern: "/containers/.+/archive", method: "PUT", action: ActionContainerArchiveExtract},
	{pattern: "/containers/.+/archive", method: "HEAD", action: ActionContainerArchiveInfo},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#get-an-archive-of-a-filesystem-resource-in-a-container
	{pattern: "/containers/.+/archive", method: "GET", action: ActionContainerArchive},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#attach-to-a-container-websocket
	{pattern: "/containers/.+/attach/ws", method: "GET", action: ActionContainerAttachWs},
	// http://docs.docker.com/reference/api/docker_remote_api_v1.21/#attach-to-a-container
	{pattern: "/containers/.+/attach", method: "POST", action: ActionContainerAttach},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#list-containers
	{pattern: "/containers/json", method: "GET", action: ActionContainerList},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#inspect-a-container
	{pattern: "/containers/.+/json", method: "GET", action: ActionContainerInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#remove-a-container
	{pattern: "/containers/.+", method: "DELETE", action: ActionContainerDelete},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#rename-a-container
	{pattern: "/containers/.+/rename", method: "POST", action: ActionContainerRename},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-container-stats-based-on-resource-usage
	{pattern: "/containers/.+/stats", method: "GET", action: ActionContainerStats},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#inspect-changes-on-a-container-s-filesystem
	{pattern: "/containers/.+/changes", method: "GET", action: ActionContainerChanges},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#list-processes-running-inside-a-container
	{pattern: "/containers/.+/top", method: "GET", action: ActionContainerTop},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-container-logs
	{pattern: "/containers/.+/logs", method: "GET", action: ActionContainerLogs},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#create-a-container
	{pattern: "/containers/create", method: "POST", action: ActionContainerCreate},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-a-tarball-containing-all-images
	{pattern: "/images/.+./get", method: "GET", action: ActionImageArchive},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#search-images
	{pattern: "/images/search", method: "GET", action: ActionImagesSearch},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#tag-an-image-into-a-repository
	{pattern: "/images/.+/tag", method: "POST", action: ActionImageTag},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#inspect-an-image
	{pattern: "/images/.+/json", method: "GET", action: ActionImageInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.18/#inspect-an-image
	{pattern: "/images/.+", method: "DELETE", action: ActionImageDelete},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#get-the-history-of-an-image
	{pattern: "/images/.+/history", method: "GET", action: ActionImageHistory},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#push-an-image-on-the-registry
	{pattern: "/images/.+/push", method: "POST", action: ActionImagePush},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#create-an-image
	{pattern: "/images/create", method: "POST", action: ActionImageCreate},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#load-a-tarball-with-a-set-of-images-and-tags-into-docker
	{pattern: "/images/load", method: "POST", action: ActionImageLoad},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#list-images
	{pattern: "/images/json", method: "GET", action: ActionImageList},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#ping-the-docker-server
	{pattern: "/_ping", method: "GET", action: ActionDockerPing},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#display-system-wide-information
	{pattern: "/info", method: "GET", action: ActionDockerInfo},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-inspect
	{pattern: "/exec/.+/json", method: "GET", action: ActionContainerExecInspect},
	// https://docs.docker.com/reference/api/docker_remote_api_v1.21/#exec-start
	{pattern: "/exec/.+/start", method: "POST", action: ActionContainerExecStart},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#inspect-a-volume
	{pattern: "/volumes/.+", method: "GET", action: ActionVolumeInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#list-volumes
	{pattern: "/volumes", method: "GET", action: ActionVolumeList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-volume
	{pattern: "/volumes/create", method: "POST", action: ActionVolumeCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#remove-a-volume
	{pattern: "/volumes/.+", method: "DELETE", action: ActionVolumeRemove},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#inspect-network
	{pattern: "/networks/.+", method: "GET", action: ActionNetworkInspect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#list-networks
	{pattern: "/networks", method: "GET", action: ActionNetworkList},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#create-a-network
	{pattern: "/networks/create", method: "POST", action: ActionNetworkCreate},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#connect-a-container-to-a-network
	{pattern: "/networks/.+/connect", method: "POST", action: ActionNetworkConnect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#disconnect-a-container-from-a-network
	{pattern: "/networks/.+/disconnect", method: "POST", action: ActionNetworkDisconnect},
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.21/#remove-a-network
	{pattern: "/networks/.+", method: "DELETE", action: ActionNetworkRemove},
}

// ParseRoute convert a method/url pattern to corresponding docker action
func ParseRoute(method, url string) string {
	for _, route := range routes {
		if route.method == method {
			match, err := regexp.MatchString(route.pattern, url)
			if err == nil && match {
				return route.action
			}

		}
	}

	return ActionNone
}
