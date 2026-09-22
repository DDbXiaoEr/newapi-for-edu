/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
import React, { useCallback, useEffect, useMemo, useState } from 'react';
import {
  Button,
  Select,
  Space,
  Spin,
  Table,
  Tag,
  Typography,
} from '@douyinfe/semi-ui';
import { useTranslation } from 'react-i18next';
import { API, showError, showSuccess } from '../../../helpers';

const { Text } = Typography;

const CHANNEL_STATUS_ENABLED = 1;

export default function GroupChannelBinding() {
  const { t } = useTranslation();
  const [loading, setLoading] = useState(false);
  const [saving, setSaving] = useState(false);
  const [groups, setGroups] = useState([]);
  const [channels, setChannels] = useState([]);
  const [groupChannels, setGroupChannels] = useState({});
  const [groupModels, setGroupModels] = useState({});
  const [group, setGroup] = useState('');
  const [selected, setSelected] = useState([]);

  const applyPayload = useCallback((data) => {
    if (!data) return null;
    setGroups(data.groups || []);
    setChannels(data.channels || []);
    setGroupChannels(data.group_channels || {});
    setGroupModels(data.group_models || {});
    return data;
  }, []);

  const load = useCallback(
    async (keepGroup) => {
      setLoading(true);
      try {
        const res = await API.get('/api/group/channels');
        const data = applyPayload(res.data?.data);
        const current =
          keepGroup && data?.groups?.includes(group)
            ? group
            : data?.groups?.[0] || '';
        setGroup(current);
        setSelected(data?.group_channels?.[current] || []);
      } catch (error) {
        showError(t('加载分组渠道绑定失败'));
      } finally {
        setLoading(false);
      }
    },
    [applyPayload, group, t],
  );

  useEffect(() => {
    load(false);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleGroupChange = (value) => {
    setGroup(value);
    setSelected(groupChannels[value] || []);
  };

  const channelById = useMemo(
    () => new Map(channels.map((item) => [item.id, item])),
    [channels],
  );

  const previewModels = useMemo(() => {
    const models = new Set();
    selected.forEach((id) => {
      const channel = channelById.get(id);
      if (!channel || channel.status !== CHANNEL_STATUS_ENABLED) return;
      (channel.models || []).forEach((model) => models.add(model));
    });
    return Array.from(models).sort();
  }, [selected, channelById]);

  const currentModels = groupModels[group] || [];
  const lostModels = useMemo(
    () => currentModels.filter((model) => !previewModels.includes(model)),
    [currentModels, previewModels],
  );

  const persist = async (ids) => {
    if (!group) return;
    setSaving(true);
    try {
      const res = await API.put('/api/group/channels', {
        group,
        channel_ids: ids,
      });
      if (!res.data?.success) {
        showError(res.data?.message || t('保存失败'));
        return;
      }
      showSuccess(
        ids.length === 0
          ? t('已取消该分组的渠道绑定')
          : t('分组渠道绑定已保存'),
      );
      await load(true);
    } catch (error) {
      showError(t('保存失败'));
    } finally {
      setSaving(false);
    }
  };

  const columns = [
    {
      title: t('渠道'),
      dataIndex: 'name',
      render: (text, record) => (
        <Text>
          {text} #{record.id}
        </Text>
      ),
    },
    {
      title: t('状态'),
      dataIndex: 'status',
      width: 110,
      render: (status) =>
        status === CHANNEL_STATUS_ENABLED ? (
          <Tag color='green'>{t('启用')}</Tag>
        ) : (
          <Tag color='grey'>{t('已禁用')}</Tag>
        ),
    },
    {
      title: t('模型数'),
      dataIndex: 'models',
      width: 100,
      render: (models) => (models || []).length,
    },
  ];

  return (
    <Spin spinning={loading}>
      <Space vertical align='start' spacing={12} style={{ width: '100%' }}>
        <Text type='tertiary'>
          {t(
            '将用户分组绑定到指定渠道后，该分组内的密钥使用模型时只会走绑定的渠道；未绑定则沿用渠道自身的分组设置。',
          )}
        </Text>
        <Space>
          <Text strong>{t('用户分组')}</Text>
          <Select
            value={group}
            onChange={handleGroupChange}
            style={{ width: 200 }}
            placeholder={t('选择分组')}
          >
            {groups.map((name) => (
              <Select.Option key={name} value={name}>
                {name}
              </Select.Option>
            ))}
          </Select>
        </Space>
        <Table
          rowKey='id'
          columns={columns}
          dataSource={channels}
          pagination={{ pageSize: 8 }}
          rowSelection={{
            selectedRowKeys: selected,
            onChange: (keys) => setSelected(keys),
            getCheckboxProps: (record) => ({
              disabled: record.status !== CHANNEL_STATUS_ENABLED,
            }),
          }}
        />
        <div>
          <Text strong>
            {t('绑定后可用模型')}（{previewModels.length}）
          </Text>
          <div style={{ marginTop: 6 }}>
            {previewModels.map((model) => (
              <Tag
                key={model}
                color='blue'
                style={{ marginRight: 6, marginBottom: 6 }}
              >
                {model}
              </Tag>
            ))}
          </div>
        </div>
        <div>
          <Text strong type={lostModels.length ? 'danger' : undefined}>
            {t('绑定后将会丢失的模型')}（{lostModels.length}）
          </Text>
          <div style={{ marginTop: 6 }}>
            {lostModels.map((model) => (
              <Tag
                key={model}
                color='red'
                style={{ marginRight: 6, marginBottom: 6 }}
              >
                {model}
              </Tag>
            ))}
          </div>
        </div>
        <Space>
          <Button
            theme='solid'
            loading={saving}
            disabled={!group}
            onClick={() => persist(selected)}
          >
            {t('保存绑定')}
          </Button>
          <Button
            loading={saving}
            disabled={!group}
            onClick={() => persist([])}
          >
            {t('取消绑定')}
          </Button>
          {selected.length === 0 && (
            <Text type='tertiary' size='small'>
              {t('未选择渠道，保存后该分组将取消绑定')}
            </Text>
          )}
        </Space>
      </Space>
    </Spin>
  );
}
