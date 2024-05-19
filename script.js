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
import { generateItUserNip, generateNurseUserNip } from './types/user.js';
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
        target: 10000,
        iterations: 1200
    });
}

export const options = {
    stages: stages,
};

export function setup() {

}
export function teardown() {

}

const client = new grpc.Client();
client.load([], 'backend.proto');

function GetItNip() {
    client.connect('127.0.0.1:50051', {
        plaintext: true
    });
    const response = client.invoke('pb.NIPService/GetItNip', {});
    client.close();
    return response.message
}
function GetNurseNip() {
    client.connect('127.0.0.1:50051', {
        plaintext: true
    });
    const response = client.invoke('pb.NIPService/GetNurseNip', {});
    client.close();
    return response.message
}

export default function () {
    let tags = {}
    if (config.LOAD_TEST) {
        const usr = TestRegister(config, GetItNip(), tags)
        TestLogin(usr, config, tags)
    } else {
        let usr
        for (let index = 0; index < 5; index++) {
            usr = TestRegister(config, generateItUserNip(), tags)
        }
        if (usr) {
            TestLogin(usr, config, generateItUserNip(), tags)
            TestNurseManagementPost(config, usr, generateNurseUserNip(), tags)
            TestNurseManagementGet(config, usr, tags)
            TestNurseManagementPut(config, usr, generateNurseUserNip(), tags)
            TestNurseManagementDelete(config, usr, tags)
            const positiveConfig = combine(config, {
                POSITIVE_CASE: true
            })
            const rawNurse = TestNurseManagementPost(positiveConfig, usr, generateNurseUserNip(), tags)
            const accessNurse = TestNurseManagementAccessPost(config, usr, rawNurse, tags)
            const nurse = TestNurseManagementLoginPost(config, accessNurse, tags)
            if (nurse) {
                TestMedicalPatientPost(config, usr, nurse, tags)
                TestMedicalPatientGet(config, usr, nurse, tags)
                TestMedicalRecordPost(config, usr, nurse, tags)
                TestMedicalRecordGet(config, usr, nurse, tags)
                TestUpload(config, usr, nurse, tags)
            }
        }
        sleep(1)
    }
}
