import React from 'react';
import { remote } from 'electron';

const FilePicker: React.FC<FilePickerProps> = ({ onSelectFile, selectedFilePath }: FilePickerProps) => {
    const openFileSelector = (): void => {
        const filePromise = remote.dialog.showOpenDialog({ properties: ['openFile'] });
        Promise.resolve(filePromise).then((res): void => {
            if (!res.canceled) {
                onSelectFile(res.filePaths[0]);
            }
        });
    };
    const handleChangeFilePath = (event: React.ChangeEvent<HTMLInputElement>): void => {
        onSelectFile(event.target.value);
    };
    return (
        <div>
            <input placeholder="No file chosen" value={selectedFilePath} onChange={handleChangeFilePath} />
            <button onClick={openFileSelector}>Browse</button>
        </div>
    );
};

type FilePickerProps = {
    onSelectFile: (filePath: string) => void;
    selectedFilePath: string;
};
export default FilePicker;
