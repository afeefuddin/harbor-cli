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
package root

import (
	"github.com/spf13/cobra"
	"github.com/goharbor/harbor-cli/pkg/api"
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/views/auditlogs"
)


func AuditLogs() *cobra.Command {

	var opts api.ListFlags

	cmd := &cobra.Command{
		Use:     "auditlogs",
		Short:   "list auditlogs",
		Long:    `Manage audit logs in Harbor repository.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.PageSize > 100 {
				return fmt.Errorf("page size should be less than or equal to 100")
			}

			auditLog, err := api.ListAuditLog()

			if err != nil {
				return fmt.Errorf("failed to get audit logs: %v", err)
			}

			auditlogs.PrintAuditLogs(auditLog)
			return nil
		},
	}
	return cmd
}
