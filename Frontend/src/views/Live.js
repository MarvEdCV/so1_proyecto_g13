import { Card, CardHeader, CardBody, CardTitle, CardSubtitle,  Row, Col, Label} from "reactstrap"
import Chart from 'react-apexcharts'
import Select from 'react-select'

import { selectThemeColors } from '@utils'

import '@styles/react/libs/charts/apex-charts.scss'
import '@styles/react/libs/flatpickr/flatpickr.scss'


const partidos = [
  { value: '1', label: '4, 1/2, Final' },
  { value: '2', label: '1, 1/8, Octavos' },
  { value: '3', label: '3, 1/4, Cuartos' },
  { value: '4', label: '4, 1/2, Final' },
  { value: '5', label: '5, 1/8 Octavos' }
]

const partidosPaises = [
  { value: '1', label:'Colombia - Mexico' },
  { value: '2', label: 'Brazil - Alemania' },
  { value: '3', label: 'Estados Unidos - Japon' },
  { value: '4', label: 'Rusia - España' },
  { value: '5', label: 'Argentina - Colombia' }
]
const options = {
  chart: {
    parentHeightOffset: 0,
    toolbar: {
      show: false
    }
  },
  plotOptions: {
    bar: {
      horizontal: true,
      barHeight: '30%',
      borderRadius: [10]
    }
  },
  grid: {
    xaxis: {
      lines: {
        show: false
      }
    }
  },
  colors: '#e80000',
  dataLabels: {
    enabled: false
  },
  xaxis: {
    categories: ['Brazil - Colombia', 'Alemania - Rusia', 'Inglaterra - Mexico', 'Argentina - España']
  },
  yaxis: {
    opposite:true
  }
}

// ** Chart Series
const series = [
  {
    data: ['5 - 1', '4 - 2', '3 - 1', '2 - 0']
  }
]

const Live = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Predicciones de los Fans 🙌</CardTitle>
      </CardHeader>
      <CardBody>
      <Row>
          <Col className='mb-1' md='6' sm='12'>
            <Label className='form-label'>Status Partidos</Label>
            <Select
              theme={selectThemeColors}
              className='react-select'
              classNamePrefix='select'
              defaultValue={partidos[0]}
              options={partidos}
              isClearable={false}
            />
          </Col>
          <Col className='mb-1' md='6' sm='12'>
            <Label className='form-label'>Partidos</Label>
            <Select
              theme={selectThemeColors}
              className='react-select'
              classNamePrefix='select'
              defaultValue={partidosPaises[0]}
              options={partidosPaises}
              isClearable={false}
            />
          </Col>
        </Row>
        <Row className='mx-0'>
        <Col  md='8' xs='15'>
        <Chart options={options} series={series} type='bar' height={400} width={1050}/>
        </Col>
        </Row>
      </CardBody>
    </Card>
  )
}

export default Live
