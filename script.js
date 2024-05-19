/* eslint-disable no-undef */
import { sleep } from 'k6';
import { config } from './config.js';
import { combine, } from './helpers/generator.js';
import exec from 'k6/execution';
import { TestLogin } from './testCases/authLogin.js';
import { TestRegister } from './testCases/authRegister.js';
import { TestNurseManagementPost } from './testCases/nurseManagementPost.js';
import { TestNurseManagementGet } from './testCases/nurseManagementGet.js';
import { TestNurseManagementPut } from './testCases/nurseManagementPut.js';
import { TestNurseManagementDelete } from './testCases/nurseManagementDelete.js';
import { TestMedicalPatientGet } from './testCases/medicalPatientGet.js';
import { TestMedicalPatientPost } from './testCases/medicalPatientPost.js';
import { TestMedicalRecordPost } from './testCases/medicalRecordPost.js';
import { TestNurseManagementAccessPost } from './testCases/nurseManagementAccessPost.js';
import { TestNurseManagementLoginPost } from './testCases/nurseManagementLoginPost.js';
import { TestUpload } from './testCases/uploadPost.js';
import { generateItUserNip } from './types/user.js';
import { TestMedicalRecordGet } from './testCases/medicalRecordGet.js';
import grpc from 'k6/net/grpc';

const stages = []

if (config.LOAD_TEST) {
    stages.push(
        { target: 50, iterations: 1, duration: "5s" },
        { target: 100, iterations: 1, duration: "10s" },
        { target: 150, iterations: 1, duration: "20s" },
        { target: 200, iterations: 1, duration: "20s" },
        { target: 250, iterations: 1, duration: "20s" },
        { target: 300, iterations: 1, duration: "20s" },
        { target: 600, iterations: 1, duration: "20s" },
        { target: 1200, iterations: 1, duration: "20s" },
    );
} else {
    stages.push({
        target: 1,
        iterations: 1
    });
}

// function determineStage() {
//     let elapsedTime = (exec.instance.currentTestRunDuration / 1000).toFixed(0);
//     if (elapsedTime < 5) return 1;
//     if (elapsedTime < 15) return 2;
//     if (elapsedTime < 35) return 3;
//     if (elapsedTime < 55) return 4;
//     if (elapsedTime < 75) return 5;
//     if (elapsedTime < 95) return 6;
//     if (elapsedTime < 115) return 7;
//     if (elapsedTime < 135) return 8;
//     return 9; // Remaining time
// }

export const options = {
    stages,
};

// let dbClient
// if (config.LOAD_TEST) {
//     dbClient = sql.open("postgres", "postgresql://postgres:postgres@localhost:5432/k6?sslmode=disable");
// }
// function setupLoadTest() {
//     dbClient.exec("DROP TABLE IF EXISTS test");
//     console.log('inserting 12000 nips for load test');
//     dbClient.exec(`CREATE TABLE IF NOT EXISTS test 
//     (
//         nip int8 PRIMARY KEY, 
//         is_used BOOLEAN DEFAULT FALSE NOT NULL,
//         is_it BOOLEAN NOT NULL,
//         access_token TEXT
//     )`);
//     dbClient.exec(`CREATE INDEX IF NOT EXISTS is_used_index ON test (is_used)`);
//     dbClient.exec(`CREATE INDEX IF NOT EXISTS is_it_index ON test (is_it)`);

//     let count = 0;
//     generateNIPs('615', (nip) => {
//         dbClient.exec(`insert into test (nip, is_it) values (${nip}, true)`);
//         count++;
//         if (count % 1000 === 0) {
//             console.log(`Inserted ${count} rows`);
//         }
//     })
//     generateNIPs('303', (nip) => {
//         dbClient.exec(`insert into test (nip, is_it) values (${nip}, false)`);
//         count++;
//         if (count % 1000 === 0) {
//             console.log(`Inserted ${count} rows`);
//         }
//     })
// }

export function setup() {
    // if (config.LOAD_TEST) {
    //     setupLoadTest()
    // }
}
export function teardown() {
    // if (config.LOAD_TEST) {
    //     // dbClient.exec("DELETE FROM test")
    //     dbClient.close()
    // }
}

// function getUnusedItNip(operationFn) {
//     dbClient.query('BEGIN'); // Begin transaction

//     const query = `WITH segmented_nips AS (
//     SELECT nip, ntile(100) OVER (ORDER BY nip) AS segment
//     FROM test
//     WHERE is_it = true
// )
// SELECT nip
// FROM segmented_nips
// WHERE segment = $segmentNumber
//     AND nip IN (
//         SELECT nip FROM test WHERE is_used = false AND is_it = true
//     )
// FOR UPDATE SKIP LOCKED
// LIMIT 1;`;
// const res = dbClient.query(query);
// if (!res || res.length === 0) {
//     dbClient.query('ROLLBACK'); // Rollback transaction if no result
//     return null;
// }
// const nip = res[0].nip;

// const opRes = operationFn(nip)
// // if (!opRes) {
// dbClient.query('ROLLBACK'); // Rollback transaction if operation fails
// return null;
// // }

// const updateQuery = `UPDATE test SET is_used = true WHERE nip = ${nip}`;
// dbClient.query(updateQuery);

// dbClient.query('COMMIT'); // Commit transaction
// return nip;
// }

const client = new grpc.Client();
client.load([''], 'backend.proto');

export default function () {
    client.connect('localhost:50051', {});
    const data = {};
    const response = client.invoke('pb.NIPService/GetItNip', data);
    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    // let tags = {}

    // if (config.LOAD_TEST) {
    //     getUnusedItNip((nip) => {
    //         console.log(nip)
    //         return false;
    //         // const usr = TestRegister(config, tags)
    //     })
    // } else {
    // let usr
    // for (let index = 0; index < 5; index++) {
    //     usr = TestRegister(config, generateItUserNip(), tags)
    // }
    // if (usr) {
    //     TestLogin(usr, config, tags)
    //     TestNurseManagementPost(config, usr, tags)
    //     TestNurseManagementGet(config, usr, tags)
    //     TestNurseManagementPut(config, usr, tags)
    //     TestNurseManagementDelete(config, usr, tags)
    //     const positiveConfig = combine(config, {
    //         POSITIVE_CASE: true
    //     })
    //     const rawNurse = TestNurseManagementPost(positiveConfig, usr, tags)
    //     const accessNurse = TestNurseManagementAccessPost(config, usr, rawNurse, tags)
    //     const nurse = TestNurseManagementLoginPost(config, accessNurse, tags)
    //     if (nurse) {
    //         TestMedicalPatientPost(config, usr, nurse, tags)
    //         TestMedicalPatientGet(config, usr, nurse, tags)
    //         TestMedicalRecordPost(config, usr, nurse, tags)
    //         TestMedicalRecordGet(config, usr, nurse, tags)
    //         TestUpload(config, usr, nurse, tags)
    //     }
    // }
    sleep(1)
    // }
}
