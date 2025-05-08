// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package auditlogs

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/auditlog"
)

// PrintAuditLogs renders the audit logs in a table format
func PrintAuditLogs(auditLogs *auditlog.ListAuditLogsOK) {
	// Define table columns
	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Username", Width: 20},
		{Title: "Resource", Width: 20},
		{Title: "ResourceType", Width: 20},
		{Title: "Operation", Width: 25},
	}

	// Populate table rows with audit log data
	var rows []table.Row
	for _, log := range auditLogs.Payload {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", log.ID),
			log.Username,
			log.Resource,
			log.ResourceType,
			log.Operation,
		})
	}

	// Create table model
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	// Render table
	fmt.Println(t.View())
}
