import {Button, Badge, Card} from 'react-bootstrap';

function MyCard(props) {
    return (<Card style={{width: '18rem', display: "inline-block"}} className="g-3">
            <Card.Header>{props.title}</Card.Header>
            <Card.Body>
                {/* <Card.Subtitle className="mb-2 text-muted">Card Subtitle</Card.Subtitle>*/}
                <Card.Text>
                    {props.text}
                </Card.Text>
                <Button variant="primary">{props.buttonText}</Button>{'  '}
                <Button variant="secondary">{props.buttonText}</Button>
            </Card.Body>
            <Card.Footer>
               <Badge bg="success">Running</Badge>
            </Card.Footer>
        </Card>);
}

export default MyCard;