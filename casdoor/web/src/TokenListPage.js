// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from "react";
import {Link} from "react-router-dom";
import {Button, Col, Popconfirm, Row, Table} from 'antd';
import moment from "moment";
import * as Setting from "./Setting";
import * as TokenBackend from "./backend/TokenBackend";
import i18next from "i18next";

class TokenListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      tokens: null,
    };
  }

  UNSAFE_componentWillMount() {
    this.getTokens();
  }

  getTokens() {
    TokenBackend.getTokens("admin")
      .then((res) => {
        this.setState({
          tokens: res,
        });
      });
  }

  newToken() {
    return {
      owner: "admin", // this.props.account.tokenname,
      name: `token_${this.state.tokens.length}`,
      createdTime: moment().format(),
      application: "app-built-in",
      accessToken: "",
      expiresIn: 7200,
      scope: "read",
      tokenType: "Bearer",
    }
  }

  addToken() {
    const newToken = this.newToken();
    TokenBackend.addToken(newToken)
      .then((res) => {
          Setting.showMessage("success", `Token added successfully`);
          this.setState({
            tokens: Setting.prependRow(this.state.tokens, newToken),
          });
        }
      )
      .catch(error => {
        Setting.showMessage("error", `Token failed to add: ${error}`);
      });
  }

  deleteToken(i) {
    TokenBackend.deleteToken(this.state.tokens[i])
      .then((res) => {
          Setting.showMessage("success", `Token deleted successfully`);
          this.setState({
            tokens: Setting.deleteRow(this.state.tokens, i),
          });
        }
      )
      .catch(error => {
        Setting.showMessage("error", `Token failed to delete: ${error}`);
      });
  }

  renderTable(tokens) {
    const columns = [
      {
        title: i18next.t("general:Name"),
        dataIndex: 'name',
        key: 'name',
        width: '300px',
        sorter: (a, b) => a.name.localeCompare(b.name),
        render: (text, record, index) => {
          return (
            <Link to={`/tokens/${text}`}>
              {text}
            </Link>
          )
        }
      },
      {
        title: i18next.t("general:Created time"),
        dataIndex: 'createdTime',
        key: 'createdTime',
        width: '160px',
        sorter: (a, b) => a.createdTime.localeCompare(b.createdTime),
        render: (text, record, index) => {
          return Setting.getFormattedDate(text);
        }
      },
      {
        title: i18next.t("general:Application"),
        dataIndex: 'application',
        key: 'application',
        width: '120px',
        sorter: (a, b) => a.application.localeCompare(b.application),
        render: (text, record, index) => {
          return (
            <Link to={`/applications/${text}`}>
              {text}
            </Link>
          )
        }
      },
      {
        title: i18next.t("general:Organization"),
        dataIndex: 'organization',
        key: 'organization',
        width: '120px',
        sorter: (a, b) => a.organization.localeCompare(b.organization),
        render: (text, record, index) => {
          return (
            <Link to={`/organizations/${text}`}>
              {text}
            </Link>
          )
        }
      },
      {
        title: i18next.t("general:User"),
        dataIndex: 'user',
        key: 'user',
        width: '120px',
        sorter: (a, b) => a.user.localeCompare(b.user),
        render: (text, record, index) => {
          return (
            <Link to={`/users/${record.organization}/${text}`}>
              {text}
            </Link>
          )
        }
      },
      {
        title: i18next.t("general:Authorization code"),
        dataIndex: 'code',
        key: 'code',
        // width: '150px',
        sorter: (a, b) => a.code.localeCompare(b.code),
        render: (text, record, index) => {
          return Setting.getClickable(text);
        }
      },
      {
        title: i18next.t("general:Access token"),
        dataIndex: 'accessToken',
        key: 'accessToken',
        // width: '150px',
        sorter: (a, b) => a.accessToken.localeCompare(b.accessToken),
        ellipsis: true,
        render: (text, record, index) => {
          return Setting.getClickable(text);
        }
      },
      {
        title: i18next.t("general:Expires in"),
        dataIndex: 'expiresIn',
        key: 'expiresIn',
        width: '120px',
        sorter: (a, b) => a.expiresIn - b.expiresIn,
      },
      {
        title: i18next.t("general:Scope"),
        dataIndex: 'scope',
        key: 'scope',
        width: '100px',
        sorter: (a, b) => a.scope.localeCompare(b.scope),
      },
      // {
      //   title: i18next.t("token:Token type"),
      //   dataIndex: 'tokenType',
      //   key: 'tokenType',
      //   width: '130px',
      //   sorter: (a, b) => a.tokenType.localeCompare(b.tokenType),
      // },
      {
        title: i18next.t("general:Action"),
        dataIndex: '',
        key: 'op',
        width: '170px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{marginTop: '10px', marginBottom: '10px', marginRight: '10px'}} type="primary" onClick={() => this.props.history.push(`/tokens/${record.name}`)}>{i18next.t("general:Edit")}</Button>
              <Popconfirm
                title={`Sure to delete token: ${record.name} ?`}
                onConfirm={() => this.deleteToken(index)}
              >
                <Button style={{marginBottom: '10px'}} type="danger">{i18next.t("general:Delete")}</Button>
              </Popconfirm>
            </div>
          )
        }
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={tokens} rowKey="name" size="middle" bordered pagination={{pageSize: 100}}
               title={() => (
                 <div>
                   {i18next.t("general:Tokens")}&nbsp;&nbsp;&nbsp;&nbsp;
                   <Button type="primary" size="small" onClick={this.addToken.bind(this)}>{i18next.t("general:Add")}</Button>
                 </div>
               )}
               loading={tokens === null}
        />
      </div>
    );
  }

  render() {
    return (
      <div>
        <Row style={{width: "100%"}}>
          <Col span={1}>
          </Col>
          <Col span={22}>
            {
              this.renderTable(this.state.tokens)
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default TokenListPage;
