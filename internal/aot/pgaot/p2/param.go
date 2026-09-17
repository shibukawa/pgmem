package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bind_param_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 < v2 {
		m.G0 = v7 + int32(80)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v12 != 0 {
			v14 = v7 - int32(-64)
			F_initStringInfo(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_bind_param_error_callback[0]))
				F_appendStringInfoStringQuoted(m, v14, v17, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+64))
					v23 = v22
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v24 == int32(0) {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v57 = v55 + int32(1)
							if v23 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
								F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_0), v7)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									m.G0 = v7 + int32(80)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v23
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
								F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_1), v7+int32(16))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_pfree(m, v23)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										m.G0 = v7 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
						if v27 == int32(0) {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v57 = v55 + int32(1)
								if v23 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
									F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_0), v7)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										m.G0 = v7 + int32(80)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v23
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
									F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_1), v7+int32(16))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											m.G0 = v7 + int32(80)
											return
										}
									}
								}
							}
						} else {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v35 = v33 + int32(1)
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v23 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v23
									*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
									*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v36
									F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_2), v7+int32(48))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											m.G0 = v7 + int32(80)
											return
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v35
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v36
									F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_3), v7+int32(32))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										m.G0 = v7 + int32(80)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = v2
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v24 == int32(0) {
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v57 = v55 + int32(1)
					if v23 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
						F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_0), v7)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							m.G0 = v7 + int32(80)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v23
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
						F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_1), v7+int32(16))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_pfree(m, v23)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v7 + int32(80)
								return
							}
						}
					}
				}
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				if v27 == int32(0) {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v57 = v55 + int32(1)
						if v23 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
							F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_0), v7)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v7 + int32(80)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v23
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
							F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_1), v7+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_pfree(m, v23)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									m.G0 = v7 + int32(80)
									return
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v35 = v33 + int32(1)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v23 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v23
							*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v36
							F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_2), v7+int32(48))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_pfree(m, v23)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									m.G0 = v7 + int32(80)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v36
							F_errcontext_msg(m, int32(_a_F_bind_param_error_callback_3), v7+int32(32))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								m.G0 = v7 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_generate_new_exec_param(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v8 = F_palloc0(m, int32(28))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(4294967304)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v18 = v16
		} else {
			v18 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
		v22 = F_lappend_oid(m, v21, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
			return v8
		}
	}
}
