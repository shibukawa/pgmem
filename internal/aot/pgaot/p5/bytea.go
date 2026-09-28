package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bytea_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v4 = int32(_a_F_bytea_sortsupport_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0])) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(1387)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v12 == int32(1) {
		v16 = F_palloc(m, int32(56))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(4596373779694328218)
			F_initHyperLogLog(m, v16, int32(10))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				F_initHyperLogLog(m, v16+int32(24), int32(10))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
					*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(1388)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(1389)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(118)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v35
					*(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0])) = v5
					return int64(0)
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0])) = v5
		return int64(0)
	}
}
func F_bytea_uuid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v95 int64
	_ = v95
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v16 == int32(1) {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v48 = int32(4)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v51 = F_errsave_start(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					if v51 == int32(0) {
						v95 = v6
						m.G0 = v9 + int32(32)
						return v95
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_bytea_uuid_0)
							F_errmsg(m, int32(_a_F_bytea_uuid_1), v9+int32(16))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v48
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(16)
								v69 = F_errdetail(m, int32(_a_F_bytea_uuid_2), v9)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v50, int32(_a_F_bytea_uuid_3), int32(1355), int32(_a_F_bytea_uuid_4))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int64(0)
									} else {
										v95 = v6
										m.G0 = v9 + int32(32)
										return v95
									}
								}
							}
						}
					}
				}
			} else {
				if v20 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				v45 = v31
				if v45 == int32(16) {
					v77 = F_palloc(m, int32(16))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int64(0)
					} else {
						v79 = int32(1)
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v81&v79 != 0 {
							v84 = v79
						} else {
							v84 = int32(4)
						}
						v85 = v12 + v84
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v86
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						*(*int64)(unsafe.Add(mBase, uint32(v77))) = v88
						v95 = base.I64_extend_i32_u(v77)
						m.G0 = v9 + int32(32)
						return v95
					}
				} else {
					v48 = v45
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v51 = F_errsave_start(m, v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						if v51 == int32(0) {
							v95 = v6
							m.G0 = v9 + int32(32)
							return v95
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_bytea_uuid_0)
								F_errmsg(m, int32(_a_F_bytea_uuid_1), v9+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v48
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(16)
									v69 = F_errdetail(m, int32(_a_F_bytea_uuid_2), v9)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										F_errsave_finish(m, v50, int32(_a_F_bytea_uuid_3), int32(1355), int32(_a_F_bytea_uuid_4))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int64(0)
										} else {
											v95 = v6
											m.G0 = v9 + int32(32)
											return v95
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v32 = int32(1)
			if v16&v32 != 0 {
				v45 = int32(base.Ui32(v16)>>(uint(v32)%32)) - v32
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v45 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
			}
			if v45 == int32(16) {
				v77 = F_palloc(m, int32(16))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					v79 = int32(1)
					v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					if v81&v79 != 0 {
						v84 = v79
					} else {
						v84 = int32(4)
					}
					v85 = v12 + v84
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v86
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
					*(*int64)(unsafe.Add(mBase, uint32(v77))) = v88
					v95 = base.I64_extend_i32_u(v77)
					m.G0 = v9 + int32(32)
					return v95
				}
			} else {
				v48 = v45
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v51 = F_errsave_start(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					if v51 == int32(0) {
						v95 = v6
						m.G0 = v9 + int32(32)
						return v95
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_bytea_uuid_0)
							F_errmsg(m, int32(_a_F_bytea_uuid_1), v9+int32(16))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v48
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(16)
								v69 = F_errdetail(m, int32(_a_F_bytea_uuid_2), v9)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v50, int32(_a_F_bytea_uuid_3), int32(1355), int32(_a_F_bytea_uuid_4))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int64(0)
									} else {
										v95 = v6
										m.G0 = v9 + int32(32)
										return v95
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
