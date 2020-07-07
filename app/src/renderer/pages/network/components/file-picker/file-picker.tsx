import React from 'react';
import { remote } from 'electron';

import Input from '../../../../components/input/input';
import Button from '../../../../components/button/button';

const FilePicker: React.FC<FilePickerProps> = ({ onSelectFile, selectedFilePath }: FilePickerProps) => {
    const openFileSelector = (): void => {
        const filePromise = remote.dialog.showOpenDialog({ properties: ['openFile'] });
        Promise.resolve(filePromise).then((res): void => {
            if (!res.canceled) {
                onSelectFile(res.filePaths[0]);
            }
        });
    };
    const handleChangeFilePath = (value: string): void => {
        onSelectFile(value);
    };
    return (
        <div style={{display: 'flex', flexDirection: 'row', alignItems: 'center'}}>
            {/* @ts-ignore */}
            <Input placeholder="No file chosen" value={selectedFilePath} onChange={handleChangeFilePath} />
            <div style={{marginLeft: '8px'}}>
                <Button onClick={openFileSelector}>Browse</Button>
            </div>
        </div>
    );
};

type FilePickerProps = {
    onSelectFile: (filePath: string) => void;
    selectedFilePath: string;
};
export default FilePicker;
