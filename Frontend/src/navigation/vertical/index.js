import { Monitor, FileText } from "react-feather"

export default [
  {
    id: "Live",
    title: "Live",
    icon: <Monitor size={20} />,
    navLink: "/Live"
  },
  {
    id: "Logs",
    title: "Logs",
    icon: <FileText size={20} />,
    navLink: "/Logs"
  }
]
