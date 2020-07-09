import React, { Component } from 'react';

import './title.scss';

type TitlePropTypes = {
    children: React.ReactNode,
};

export default class Title extends Component<TitlePropTypes> {
    static PageTitle(props: TitlePropTypes): JSX.Element {
        return (
            <h1 className="title--page-title">{props.children}</h1>
        );
    }

    static SmallTitle(props: TitlePropTypes): JSX.Element {
        return (
            <h3 className="title--small-title">{props.children}</h3>
        );
    }

    static SectionTitle(props: TitlePropTypes): JSX.Element {
        return (
            <h5 className="title--section-title">{props.children}</h5>
        );
    }

    render(): JSX.Element {
        return (
            <h2 className="title--title">{this.props.children}</h2>
        );
    }
}
