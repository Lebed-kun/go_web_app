import React from 'react';
import { defaultHead as Head } from 'next/head';
import { Card, Col } from 'antd';

import Layout from '../components/Layout.jsx';

import { REST_BASE_URL } from '../constants.js';
import Axios from 'axios';

function Index(props) {
    let content = null;
    if (props.tasks) {
        content = props.tasks.map((el, id) => (
            <Col span={6}>
                <Card key={id} title={el.Title ? el.Title : el.Description.slice(0, 20)}>
                    <p>
                        {el.Description.slice(0, 80)}
                    </p>
                </Card>
            </Col>
        ))
    }
    
    return (
        <Layout>
            <Head>
                <title>Tasks</title>
            </Head>

            {content}
        </Layout>
    )
}

Index.getInitialProps = async ({ request }) => {
    const response = await Axios.get(`${REST_BASE_URL}`);
    const json = await response.json();
    return { tasks : json }
}

export default Index;



