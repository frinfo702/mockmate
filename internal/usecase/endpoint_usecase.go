package usecase

import "github.com/frinfo702/mockmate/internal/entity"

// SelectVersion returns the endpoint version based on the requested version (if provided)
// If queryVersion is non-empty and matches one of the versions, that version is returned.
// Otherwise, the first version is returned.
func SelectVersion(ver entity.EndPointVersion, queryVersion string) entity.EndPointVersion {
	// 単一のVersionが対象の場合、シンプルに返す実装例（拡張する場合はエンドポイント全体から選ぶ）
	if queryVersion != "" && queryVersion == ver.Version {
		return ver
	}
	// 今回はシンプルに、常に受け取った ver を返す（各ルートは個別に登録している前提）
	return ver
}
