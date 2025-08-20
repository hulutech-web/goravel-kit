declare namespace API {
  type AuthCodeRequest = {
    code?: string;
  };

  type destroyParams = {
    /** id */
    id: string;
  };

  type destroyParams = {
    /** id */
    id: string;
  };

  type destroyParams = {
    /** id */
    id: string;
  };

  type FileCateRequest = {
    name?: string;
    pid?: string;
    sort?: string;
    tenant_id?: string;
    type?: string;
  };

  type FileDelRequest = {
    del_ids?: number[];
  };

  type FileRequest = {
    cid?: number;
    engine?: string;
    ext?: string;
    name?: string;
    path?: string;
    size?: number;
    tenant_id?: number;
    type?: string;
    uri?: string;
    user_id?: string;
  };

  type HouseRequest = {
    address?: string;
    albums?: string;
    area?: string;
    deposit?: string;
    description?: string;
    facilities?: string;
    header_img?: string;
    landlord_id?: string;
    location?: string;
    monthly_rent?: string;
    poster?: string;
    property_fee?: string;
    shopping?: string;
    status?: string;
    swipers?: string;
    title?: string;
    traffic?: string;
    video?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type indexParams = {
    /** name */
    name?: string;
    /** pageSize */
    pageSize?: string;
    /** currentPage */
    currentPage?: string;
    /** sort */
    sort?: string;
    /** order */
    order?: string;
  };

  type listParams = {
    /** username */
    username?: string;
  };

  type LoginRequest = {
    account?: string;
    password?: string;
  };

  type MenuRequest = {
    badge?: string;
    cacheable?: string;
    component?: string;
    icon?: string;
    menu_type?: string;
    name?: string;
    path?: string;
    permission?: string;
    pid?: string;
    render_menu?: string;
    sort?: string;
    target?: string;
    title?: string;
  };

  type optionParams = {
    /** username */
    username?: string;
  };

  type PermissionRequest = {
    code?: string;
    description?: string;
    menu_id?: string;
    name?: string;
    type?: string;
  };

  type RoleRequest = {
    is_admin?: string;
    is_disable?: string;
    label?: string;
    name?: string;
    remark?: string;
    sort?: string;
    tenant_id?: string;
  };

  type SyncRolePerRequest = {
    formIDs?: number[];
  };

  type updateParams = {
    /** id */
    id: string;
  };

  type updateParams = {
    /** id */
    id: string;
  };

  type updateParams = {
    /** id */
    id: string;
  };

  type updateParams = {
    /** id */
    id: string;
  };

  type User = true;

  type UserRequest = {
    address?: string;
    avatar?: string;
    email?: string;
    is_disable?: string;
    is_member?: string;
    is_multipoint?: string;
    openid?: string;
    password?: string;
    phone?: string;
    remark?: string;
    role_id?: string;
    salt?: string;
    sex?: string;
    tenant_id?: string;
    unionid?: string;
    username?: string;
  };
}
