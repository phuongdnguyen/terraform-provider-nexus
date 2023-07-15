package other

func flattenCriteria(criteria *Criteria) []map[string]interface{} {
	if criteria == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"last_downloaded_days":   criteria.LastDownloaded,
			"last_blob_updated_days": criteria.LastBlobUpdated,
			"regex":                  criteria.Regex,
		},
	}
}
