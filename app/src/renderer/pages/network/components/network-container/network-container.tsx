import React, { FunctionComponent } from 'react';

import FilePicker from '../file-picker';

const NetworkContainer: FunctionComponent<NetworkContainerProps> = (props) => {
    const [selectedFilePath , setSelectedFilePath] = React.useState<string>('');
    const handleSelectFile = (filePath: string): void=> {
        setSelectedFilePath(filePath);
    };
    return (
        <div>
            <h2>Network</h2>
            <div>Select a configure file by browsing for it, or by dragging it onto screen:</div>
            <div>
                <FilePicker onSelectFile={handleSelectFile} selectedFilePath={selectedFilePath}/>
            </div>
        </div>
    );
};

type NetworkContainerProps = {

}

export default NetworkContainer;
