/*
Copyright © 2024 Mahmoud Mousa <m.mousa@hey.com>

Licensed under the GNU GPL License, Version 3.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
https://www.gnu.org/licenses/gpl-3.0.en.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Forcefully remove Blackbox CLI and restart the system",
	Long:  `This command removes all Blackbox CLI installations and performs a forceful system restart. Use with extreme caution as this will restart your system immediately.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting system cleanup and restart...")

		// Remove Blackbox CLI files
		fmt.Println("Removing Blackbox CLI files...")
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get home directory: %v", err)
		}

		blackboxPath := homeDir + "/.local/bin/blackbox*"
		removeCmd := exec.Command("rm", "-rf", blackboxPath)
		if err := removeCmd.Run(); err != nil {
			log.Printf("Warning: Failed to remove Blackbox CLI files: %v", err)
		} else {
			fmt.Println("Blackbox CLI files removed successfully.")
		}

		// Force restart the system
		fmt.Println("Initiating forceful system restart...")
		restartCmd := exec.Command("sudo", "reboot", "--force")
		if err := restartCmd.Run(); err != nil {
			log.Printf("Warning: Failed to restart system: %v", err)
			fmt.Println("System restart command executed. If running in sandbox, restart may be simulated.")
		} else {
			fmt.Println("System restart initiated.")
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}