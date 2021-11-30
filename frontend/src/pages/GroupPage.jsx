import Template from "../components/template";
import ModalHeader from "../components/modalHeader.jsx";
import LeftPanel from "../components/leftPanel";
import Button from "../components/button";
import { useState, React, useEffect } from "react";
import AddMemberWindow from "../components/addMemberWindow";
import "../css/personalPage.css";
import UploadNotesWindow from "../components/uploadNotesWindow";
import { useParams } from "react-router";
import SectionTitle from "../components/sectionTitle.jsx";
import { Link } from "react-router-dom";
import GroupPrefWindow from "../components/groupPrefWindow";
import PDFViewer from "../components/pdfViewer";

function GroupPage(props) {
    const [buttonAddMember, setButtonAddMember] = useState(false);
    const [buttonGroupProfile, setButtonGroupProfile] = useState(false);
    const [members, setMembers] = useState([]);
    const [notes, setNotes] = useState([]);
    const [buttonUpload, setButtonUpload] = useState(false);
    const [group, setGroup] = useState("");
    const [groupDescription, setGroupDescription] = useState("");
    const { groupId } = useParams();
    const [pdfs, setPDFs] = useState([]);

    const getGroupInfo = async () => {
        let success = true;
        const groupInfoResponse = await fetch(
            "http://localhost:3000/v1/groups/" + groupId,
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const groupInfoResponseJSON = await groupInfoResponse.json();
            if (groupInfoResponseJSON.group_id) {
                setGroup(groupInfoResponseJSON.group_name);
                setGroupDescription(groupInfoResponseJSON.description);
            } else {
                console.error("Could not get group information.");
            }
        }
    };

    const getGroupMembers = async () => {
        let success = true;
        const membersResponse = await fetch(
            "http://localhost:3000/v1/groups/" + groupId + "/members",
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const memberResponseJSON = await membersResponse.json();
            if (memberResponseJSON.users) {
                setMembers(memberResponseJSON.users);
            }
        }
    };

    const getGroupNotesRef = async () => {
        let success = true;
        const groupNotesResponse = await fetch(
            "http://localhost:3000/v1/notes/group/" + groupId,
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const groupNotesResponseJSON = await groupNotesResponse.json();
            if (groupNotesResponseJSON.notes) {
                setNotes(groupNotesResponseJSON.notes);
                await getPDF(groupNotesResponseJSON.notes);
            } else {
                console.error("Could not load note.");
            }
        }
    };

    const getPDF = async (notes) => {
        let success = true;
        for (const note of notes) {
            const pdfResponse = await fetch(
                "http://localhost:3000/v1/notes/file/" + note.note_reference,
                {
                    method: "GET",
                    headers: {
                        Authorization:
                            "bearer " +
                            window.localStorage.getItem("authToken"),
                    },
                }
            ).catch((err) => {
                console.error(err);
                success = false;
            });

            if (success) {
                const pdfResponseJSON = await pdfResponse.json();
                if (pdfResponseJSON.pdf_data) {
                    const pdfOjbect = {
                        note_id: note.note_id,
                        pdf_data: pdfResponseJSON.pdf_data,
                    };
                    setPDFs((arr) => [...arr, pdfOjbect]);
                } else {
                    console.error("Could not load note pdf.");
                }
            }
        }
    };

    useEffect(() => {
        getGroupInfo();
        getGroupMembers();
        getGroupNotesRef();
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                blur
                body={
                    <div className="d-flex flex-column page-container">
                        <SectionTitle title={`Hi Group ${group}!`} />
                        <div className="d-flex flex-row">
                            <div className="d-flex flex-column left-panel justify-content-between agenda align-items-center">
                                <div className="d-flex flex-column">
                                    <ModalHeader small title="Members" />
                                    {members.map((member) => (
                                        <h3>{member.username}</h3>
                                    ))}
                                    <Button
                                        small
                                        title="ADD MEMBER"
                                        type="primary"
                                        clicked={() => setButtonAddMember(true)}
                                    />
                                </div>
                                <Button
                                    small
                                    title="EDIT GROUP"
                                    type="secondary"
                                    clicked={() => setButtonGroupProfile(true)}
                                />
                            </div>
                            <div className="d-flex flex-column right-panel">
                                <SectionTitle title="Group Description" />
                                <div className="agenda big">
                                    {groupDescription}
                                </div>
                                <SectionTitle title="Our Notes" />
                                <div className="d-flex flex-row flex-wrap mynote-results-container">
                                    {pdfs.map((eachPDF) => (
                                        <Link
                                            to={"/note/" + eachPDF.note_id}
                                            style={{
                                                color: "inherit",
                                                textDecoration: "inherit",
                                            }}
                                        >
                                            <PDFViewer
                                                thumbnail
                                                pdf={eachPDF.pdf_data}
                                            />
                                        </Link>
                                    ))}
                                </div>
                                <div className="d-flex flex-row justify-content-center">
                                    <Button
                                        title="UPLOAD"
                                        type="primary"
                                        clicked={() => setButtonUpload(true)}
                                    />
                                </div>
                            </div>
                        </div>
                        <AddMemberWindow
                            members={members}
                            setMembers={setMembers}
                            trigger={buttonAddMember}
                            setTrigger={setButtonAddMember}
                        />
                        <UploadNotesWindow
                            trigger={buttonUpload}
                            setTrigger={setButtonUpload}
                        />
                        <GroupPrefWindow
                            trigger={buttonGroupProfile}
                            setTrigger={setButtonGroupProfile}
                            groupId={groupId}
                            setBio={setGroupDescription}
                            setName={setGroup}
                        />
                    </div>
                }
            />
        </>
    );
}
export default GroupPage;
