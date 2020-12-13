package commands

const indent = `&nbsp;&nbsp;&nbsp;&nbsp;`

func displayHelpMsg() string {
	return `__help__` + "   \n" +
		indent + `List all available bot commands.` + "   \n" +
		indent + `OPTIONAL` + "   \n" +
		indent + indent + `Command name - prints the command documentation` + "   \n" +
		indent + indent + `Usage: ` + "   \n" +
		indent + indent + indent + `help <command>` + "   \n" +
		indent + indent + `Example:` + "   \n" +
		indent + indent + indent + `help listPRs` + "   \n" +
		`__listPRs__` + "   \n" +
		indent + `List all active pull requests.` + "   \n" +
		indent + `OPTIONAL` + "   \n" +
		indent + indent + `Name of the reviewer - returns the pull requests asigned to <name of reviewer>` + "   \n" +
		indent + indent + `Usage:` + "   \n" +
		indent + indent + indent + `listPRs <name of reviewer>` + "   \n" +
		indent + indent + `Example:` + "   \n" +
		indent + indent + indent + `listPRs Pepito Perez` + "   \n"

}
