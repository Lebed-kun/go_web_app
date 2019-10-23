import React from 'react';
import { defaultHead as Head } from 'next/head';

import Layout from '../components/Layout.jsx';

import { REST_BASE_URL } from '../constants.js';
import Axios from 'axios';

function Index() {
    return (
        <Layout>

        </Layout>
    )
}

Index.getInitialProps = async ({ request }) => {
    const response = await Axios.get(`${REST_BASE_URL}/`);
    const json = await response.json();
    return { tasks : json }
}



