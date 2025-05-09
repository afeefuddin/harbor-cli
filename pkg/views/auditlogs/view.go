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
	"os"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/auditlog"
	"github.com/goharbor/harbor-cli/pkg/views/base/tablelist"
)

// PrintAuditLogs renders the audit logs in a table format
func PrintAuditLogs(auditLogs *auditlog.ListAuditLogsOK) {
	// Define table columns
	columns := []table.Column{
		{Title: "ID", Width: tablelist.WidthS},
		{Title: "Username", Width: tablelist.WidthXXL},
		{Title: "Resource", Width: tablelist.WidthXXL},
		{Title: "ResourceType", Width: tablelist.WidthS},
		{Title: "Operation", Width: tablelist.WidthS},
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
	m := tablelist.NewModel(columns, rows, len(rows))
	// Render table
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
