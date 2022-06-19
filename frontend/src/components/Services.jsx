import {Accordion, Badge, Card, Col, Container, Row, Table} from 'react-bootstrap';

function Services(props) {
    let groups = [];
    let index = 0;
    let activeKeys = [index];
    for (let key in props) {
        let value = props[key];
        groups.push(ServicesGroup({items: value, eventKey: index, header: key}));
        ++index;
        activeKeys.push(index);
    }
    return (
        <Accordion defaultActiveKey={activeKeys} alwaysOpen>
            {groups}
        </Accordion>
    )
}

function ServicesGroup(props) {
    let services = [];
    props.items.forEach(function (value) {
        services.push(ServiceCard(value));
    })

    return (
        <Accordion.Item eventKey={props.eventKey}>
            <Accordion.Header>{props.header}</Accordion.Header>
            <Accordion.Body>
                {services}
            </Accordion.Body>
        </Accordion.Item>
    )
}

function ServiceCard(props) {
    console.log(props);
    let img = props.item_type === 'DYNAMIC'
        ? (<svg width="12" height="12" className="rounded me-2">
            <rect width="12" height="12" style={{fill: 'blue'}}/>
        </svg>)
        : (<svg width="12" height="12" className="rounded me-2">
            <rect width="12" height="12" style={{fill: 'green'}}/>
        </svg>);

    return (<Card style={{width: '25rem', display: "inline-block", overflow: 'hidden'}} className="rounded me-4">
        <Card.Header>
            <Container>
                <Row>
                    <Col xs={1}>{img}</Col>
                    <Col xs={8}><strong className="me-auto">{props.name}</strong></Col>
                    <Col xs={1}><Badge bg="success">Running</Badge></Col>
                </Row>
            </Container>


        </Card.Header>
        <Card.Body>
            {/* <Card.Subtitle className="mb-2 text-muted">Card Subtitle</Card.Subtitle>*/}
            <Card.Text style={{'white-space': 'nowrap', 'overflow': 'hidden', 'text-overflow': 'ellipsis'}}>
                <Table responsive="sm">
                    <thead>
                    <tr>
                        <th><strong className="me-auto"><small>Docker</small></strong></th>
                        <th></th>
                        <th></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td></td>
                        <td><small>image</small></td>
                        <td><small>{props.docker_config.image}</small></td>
                    </tr>
                    <tr>
                        <td></td>
                        <td><small>container</small></td>
                        <td><small>{props.docker_config.container_name}</small></td>
                    </tr>
                    </tbody>
                </Table>
                {props.git_config && <Table responsive="sm">
                    <thead>
                    <tr>
                        <th><strong className="me-auto"><small>Git</small></strong></th>
                        <th></th>
                        <th></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td></td>
                        <td><small>repo</small></td>
                        <td><small>{props.git_config.repo}</small></td>
                    </tr>
                    </tbody>
                </Table>}

            </Card.Text>
        </Card.Body>
    </Card>);
}


export default Services;