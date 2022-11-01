import {
  Card,
  CardHeader,
  CardBody,
  CardTitle,
  CardText,
  CardLink
} from "reactstrap"

const Logs = () => {
  return (
    <div>
      <Card>
        <CardHeader>
          <CardTitle>Records</CardTitle>
        </CardHeader>
        <CardBody>
          <CardText>Registros de la bases de datos</CardText>
          <CardText>
            LOGS
          </CardText>
        </CardBody>
      </Card>
    </div>
  )
}

export default Logs
