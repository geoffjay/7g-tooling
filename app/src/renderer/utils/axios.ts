import axios from 'axios';
import fs from 'fs';

export default {
    postFile: (filePath: string, callback: (status: number) => void): void => {
        fs.readFile(filePath, 'utf8', (err, data) => {
            if (err) {
                return callback(400);
            }
            const formData = new FormData();
            formData.append('file', data);
            console.log(data);
            axios.post('http://localhost:3000/v1/network/populate', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            }).then(res => console.log(res)
            ).catch(err => console.log(err));
        });
    }
};
