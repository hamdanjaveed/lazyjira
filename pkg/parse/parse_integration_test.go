package parse_test

import (
	"strings"
	"testing"

	"github.com/hamdanjaveed/lazygit/pkg/parse"
	"github.com/stretchr/testify/assert"
)

func TestParseTable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
		err      string
	}{
		{
			// This test was generated using data from a sample Jira project by running: `jira issues list --plain --no-truncate`.
			name: "sample data",
			input: `TYPE            KEY     SUMMARY                                                                                                                                                                                                                                                 STATUS          ASSIGNEE        REPORTER        PRIORITY        RESOLUTION      CREATED                 UPDATED                 LABELS
Story           TP-9    As a developer, I'd like to update story status during the sprint >> Click the Active sprints link at the top right of the screen to go to the Active sprints where the current Sprint's items can be updated                                           To Do                           Mr. Jira        Medium                          2024-06-29 11:10:29     2024-06-29 11:10:29
Bug             TP-8    As a product owner, I'd like to include bugs, tasks and other issue types in my backlog >> Bugs like this one will also appear in your backlog but they are not normally estimated                                                                      To Do                           Mr. Jira        Medium                          2024-06-29 11:10:29     2024-06-29 11:10:29
Sub-task        TP-7    This is a sample task. Tasks are used to break down the steps to implement a user story                                                                                                                                                                 To Do           Mr. Jira        Mr. Jira        Medium                          2024-06-29 11:10:29     2024-06-29 11:10:29
Story           TP-6    As a scrum master, I'd like to break stories down into tasks we can track during the sprint >> Try creating a task by clicking the Sub-Tasks tab in the Detail View on the right                                                                        To Do                           Mr. Jira        Medium                          2024-06-29 11:10:29     2024-06-29 11:10:29
Story           TP-5    As a team, I'd like to commit to a set of stories to be completed in a sprint (or iteration) >> Click "Create Sprint" then drag the footer down to select issues for a sprint (you can't start a sprint at the moment because one is already active)    To Do                           Mr. Jira        Medium                          2024-06-29 11:10:29     2024-06-29 11:10:29
Story           TP-4    As a team, I'd like to estimate the effort of a story in Story Points so we can understand the work remaining >> Try setting the Story Points for this story in the "Estimate" field                                                                    To Do                           Mr. Jira        Medium                          2024-06-29 11:10:28     2024-06-29 11:10:28
Story           TP-3    As a product owner, I'd like to rank stories in the backlog so I can communicate the proposed implementation order >> Try dragging this story up above the previous story                                                                               To Do                           Mr. Jira        Medium                          2024-06-29 11:10:28     2024-06-29 11:10:28
Story           TP-2    As a product owner, I'd like to express work in terms of actual user problems, aka User Stories, and place them in the backlog >> Try creating a new story with the "+ Create Issue" button (top right of screen)                                       To Do                           Mr. Jira        Medium                          2024-06-29 11:10:28     2024-06-29 11:10:28
Story           TP-1    As an Agile team, I'd like to learn about Scrum >> Click the "TP-1" link at the left of this row to see detail in the Description tab on the right                                                                                                      To Do                           Mr. Jira        Medium                          2024-06-29 11:10:27     2024-06-29 11:10:27
Story           TP-14   As a user, I can find important items on the board by using the customisable "Quick Filters" above >> Try clicking the "Only My Issues" Quick Filter above                                                                                              To Do           Mr. Jira        Mr. Jira        Medium                          2024-06-28 07:00:27     2024-06-28 07:00:27
Sub-task        TP-12   When the last task is done, the story can be automatically closed >> Drag this task to "Done" too                                                                                                                                                       In Progress                     Mr. Jira        Medium                          2024-06-26 11:27:27     2024-06-26 11:27:27
Sub-task        TP-11   Update task status by dragging and dropping from column to column >> Try dragging this task to "Done"                                                                                                                                                   In Progress     Mr. Jira        Mr. Jira        Medium                          2024-06-25 14:02:27     2024-06-25 14:02:27
Bug             TP-17   Instructions for deleting this sample board and project are in the description for this issue >> Click the "TP-17" link and read the description tab of the detail view for more                                                                        Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-22 02:00:27     2024-06-25 18:36:27
Story           TP-15   As a scrum master, I can see the progress of a sprint via the Burndown Chart >> Click "Reports" to view the Burndown Chart                                                                                                                              Done                            Mr. Jira        Medium          Done            2024-06-22 02:00:27     2024-06-26 16:36:27
Bug             TP-13   As a developer, I can update details on an item using the Detail View >> Click the "TP-13" link at the top of this card to open the detail view                                                                                                         To Do           Mr. Jira        Mr. Jira        Medium                          2024-06-22 02:00:27     2024-06-22 02:00:27
Story           TP-10   As a developer, I can update story and task status with drag and drop (click the triangle at far left of this story to show sub-tasks)                                                                                                                  In Progress     Mr. Jira        Mr. Jira        Medium                          2024-06-22 02:00:27     2024-06-22 02:00:27
Story           TP-22   As a user, I'd like a historical story to show in reports                                                                                                                                                                                               Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-08 00:50:27     2024-06-18 07:47:27
Story           TP-23   As a user, I'd like a historical story to show in reports                                                                                                                                                                                               Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-08 00:50:27     2024-06-20 09:25:27
Story           TP-21   As a user, I'd like a historical story to show in reports                                                                                                                                                                                               Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-08 00:50:27     2024-06-15 16:27:27
Story           TP-18   As a user, I'd like a historical story to show in reports                                                                                                                                                                                               Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-08 00:50:27     2024-06-08 21:20:27
Story           TP-19   As a user, I'd like a historical story to show in reports                                                                                                                                                                                               Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-08 00:50:27     2024-06-11 21:14:27
Story           TP-20   As a user, I'd like a historical story to show in reports                                                                                                                                                                                               Done            Mr. Jira        Mr. Jira        Medium          Done            2024-06-08 00:50:27     2024-06-13 09:49:27
Story           TP-16   As a team, we can finish the sprint by clicking the cog icon next to the sprint name above the "To Do" column then selecting "Complete Sprint" >> Try closing this sprint now                                                                           Done                            Mr. Jira        Medium          Done            2024-06-01 05:56:27     2024-06-23 08:32:27`,
			expected: [][]string{
				{"Story", "TP-9", "As a developer, I'd like to update story status during the sprint >> Click the Active sprints link at the top right of the screen to go to the Active sprints where the current Sprint's items can be updated", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:29", "2024-06-29 11:10:29", ""},
				{"Bug", "TP-8", "As a product owner, I'd like to include bugs, tasks and other issue types in my backlog >> Bugs like this one will also appear in your backlog but they are not normally estimated", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:29", "2024-06-29 11:10:29", ""},
				{"Sub-task", "TP-7", "This is a sample task. Tasks are used to break down the steps to implement a user story", "To Do", "Mr. Jira", "Mr. Jira", "Medium", "", "2024-06-29 11:10:29", "2024-06-29 11:10:29", ""},
				{"Story", "TP-6", "As a scrum master, I'd like to break stories down into tasks we can track during the sprint >> Try creating a task by clicking the Sub-Tasks tab in the Detail View on the right", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:29", "2024-06-29 11:10:29", ""},
				{"Story", "TP-5", "As a team, I'd like to commit to a set of stories to be completed in a sprint (or iteration) >> Click \"Create Sprint\" then drag the footer down to select issues for a sprint (you can't start a sprint at the moment because one is already active)", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:29", "2024-06-29 11:10:29", ""},
				{"Story", "TP-4", "As a team, I'd like to estimate the effort of a story in Story Points so we can understand the work remaining >> Try setting the Story Points for this story in the \"Estimate\" field", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:28", "2024-06-29 11:10:28", ""},
				{"Story", "TP-3", "As a product owner, I'd like to rank stories in the backlog so I can communicate the proposed implementation order >> Try dragging this story up above the previous story", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:28", "2024-06-29 11:10:28", ""},
				{"Story", "TP-2", "As a product owner, I'd like to express work in terms of actual user problems, aka User Stories, and place them in the backlog >> Try creating a new story with the \"+ Create Issue\" button (top right of screen)", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:28", "2024-06-29 11:10:28", ""},
				{"Story", "TP-1", "As an Agile team, I'd like to learn about Scrum >> Click the \"TP-1\" link at the left of this row to see detail in the Description tab on the right", "To Do", "", "Mr. Jira", "Medium", "", "2024-06-29 11:10:27", "2024-06-29 11:10:27", ""},
				{"Story", "TP-14", "As a user, I can find important items on the board by using the customisable \"Quick Filters\" above >> Try clicking the \"Only My Issues\" Quick Filter above", "To Do", "Mr. Jira", "Mr. Jira", "Medium", "", "2024-06-28 07:00:27", "2024-06-28 07:00:27", ""},
				{"Sub-task", "TP-12", "When the last task is done, the story can be automatically closed >> Drag this task to \"Done\" too", "In Progress", "", "Mr. Jira", "Medium", "", "2024-06-26 11:27:27", "2024-06-26 11:27:27", ""},
				{"Sub-task", "TP-11", "Update task status by dragging and dropping from column to column >> Try dragging this task to \"Done\"", "In Progress", "Mr. Jira", "Mr. Jira", "Medium", "", "2024-06-25 14:02:27", "2024-06-25 14:02:27", ""},
				{"Bug", "TP-17", "Instructions for deleting this sample board and project are in the description for this issue >> Click the \"TP-17\" link and read the description tab of the detail view for more", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-22 02:00:27", "2024-06-25 18:36:27", ""},
				{"Story", "TP-15", "As a scrum master, I can see the progress of a sprint via the Burndown Chart >> Click \"Reports\" to view the Burndown Chart", "Done", "", "Mr. Jira", "Medium", "Done", "2024-06-22 02:00:27", "2024-06-26 16:36:27", ""},
				{"Bug", "TP-13", "As a developer, I can update details on an item using the Detail View >> Click the \"TP-13\" link at the top of this card to open the detail view", "To Do", "Mr. Jira", "Mr. Jira", "Medium", "", "2024-06-22 02:00:27", "2024-06-22 02:00:27", ""},
				{"Story", "TP-10", "As a developer, I can update story and task status with drag and drop (click the triangle at far left of this story to show sub-tasks)", "In Progress", "Mr. Jira", "Mr. Jira", "Medium", "", "2024-06-22 02:00:27", "2024-06-22 02:00:27", ""},
				{"Story", "TP-22", "As a user, I'd like a historical story to show in reports", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-08 00:50:27", "2024-06-18 07:47:27", ""},
				{"Story", "TP-23", "As a user, I'd like a historical story to show in reports", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-08 00:50:27", "2024-06-20 09:25:27", ""},
				{"Story", "TP-21", "As a user, I'd like a historical story to show in reports", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-08 00:50:27", "2024-06-15 16:27:27", ""},
				{"Story", "TP-18", "As a user, I'd like a historical story to show in reports", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-08 00:50:27", "2024-06-08 21:20:27", ""},
				{"Story", "TP-19", "As a user, I'd like a historical story to show in reports", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-08 00:50:27", "2024-06-11 21:14:27", ""},
				{"Story", "TP-20", "As a user, I'd like a historical story to show in reports", "Done", "Mr. Jira", "Mr. Jira", "Medium", "Done", "2024-06-08 00:50:27", "2024-06-13 09:49:27", ""},
				{"Story", "TP-16", "As a team, we can finish the sprint by clicking the cog icon next to the sprint name above the \"To Do\" column then selecting \"Complete Sprint\" >> Try closing this sprint now", "Done", "", "Mr. Jira", "Medium", "Done", "2024-06-01 05:56:27", "2024-06-23 08:32:27", ""},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := parse.ParseTable(strings.NewReader(test.input))
			if test.err != "" {
				assert.ErrorContains(t, err, test.err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, p)
		})
	}
}
