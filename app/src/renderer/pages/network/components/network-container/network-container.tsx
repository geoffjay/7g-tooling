import React, { FunctionComponent } from 'react';

import axios from '../../../../utils/axios';
import FilePicker from '../file-picker/';
import Title from '../../../../components/title/title';
import Text from '../../../../components/text/text';
import Box from '../../../../components/box/box';
import Button from '../../../../components/button/button';


const NetworkContainer: FunctionComponent<NetworkContainerProps> = (): JSX.Element => {
    const [selectedFilePath , setSelectedFilePath] = React.useState<string>('');
    const handleSelectFile = (filePath: string): void => {
        setSelectedFilePath(filePath);
    };
    const handleDeployFile = (): void => {
        axios.postFile(selectedFilePath, (status) => {
            console.log(status);
        });
    };
    return (
        <Box height="100%" flexDirection="column">
            <Box marginBottom="24px">
                <Title>Network</Title>
            </Box>
            <Text>Select a configure file by browsing for it, or by dragging it onto screen:</Text>
            <Box marginY="16px">
                <FilePicker onSelectFile={handleSelectFile} selectedFilePath={selectedFilePath}/>
            </Box>
            <Box flex={1}>
                <table/>
            </Box>
            <Box justifyContent="flex-end">
                <Button.Filled onClick={handleDeployFile}>Deploy file</Button.Filled>
            </Box>
        </Box>
    );
};

type NetworkContainerProps = {

}

export default NetworkContainer;
