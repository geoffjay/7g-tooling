import React, { FunctionComponent } from 'react';

import FilePicker from '../file-picker/';
import Title from '../../../../components/title/title';
import Text from '../../../../components/text/text';

const NetworkContainer: FunctionComponent<NetworkContainerProps> = (): JSX.Element => {
    const [selectedFilePath , setSelectedFilePath] = React.useState<string>('');
    const handleSelectFile = (filePath: string): void=> {
        setSelectedFilePath(filePath);
    };
    return (
        <div>
            <Title>Network</Title>
            <Text>Select a configure file by browsing for it, or by dragging it onto screen:</Text>
            <div>
                <FilePicker onSelectFile={handleSelectFile} selectedFilePath={selectedFilePath}/>
            </div>
        </div>
    );
};

type NetworkContainerProps = {

}

export default NetworkContainer;
