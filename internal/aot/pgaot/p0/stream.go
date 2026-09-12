package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_begin_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v12 != 0 {
		v38 = v12
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
		v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+118)))
		v42 = F_read_stream_begin_impl(m, l0, l1, l2, v38, v40, int32(0), l3, l4, l5)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v42
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = v16
		v18 = F_smgropen(m, v10, v13)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v18
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
			if v24 != 0 {
				v32 = v24
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
				v32 = v30
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v32 + int32(1)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			v38 = v36
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
			v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+118)))
			v42 = F_read_stream_begin_impl(m, l0, l1, l2, v38, v40, int32(0), l3, l4, l5)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v42
			}
		}
	}
}
func F_stream_commit_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(100710)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(4508504)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(503714)
				F_errmsg(m, int32(318986), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(497876), int32(1485), int32(218004))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _consts[77])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_stream_message_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	if v16 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(404314)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(993)
		v23 = int32(4508504)
		v24 = *(*int32)(unsafe.Add(mBase, _consts[77]))
		*(*int32)(unsafe.Add(mBase, _consts[77])) = v13 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(16)
		v33 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+147)) = uint8(v33)
		if l1 != 0 {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v37 = v36
		} else {
			v37 = int32(0)
		}
		v38 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+164)) = uint8(v38)
		*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v37
		m.T0[v16].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32))(m, v15, l1, l2, l3, l4, l5, l6)
		mBase = m.M
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		*(*int32)(unsafe.Add(mBase, _consts[77])) = v44
	} else {
	}
	m.G0 = v13 + int32(32)
	return
}
func F_stream_open_and_write_change(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	if v11 == int32(0) {
		F_stream_start_internal(m, l0, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[541]))
			v19 = v18
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v2)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21 - v22 + int32(1)
			F_BufFileWrite(m, v19, v8+int32(8), int32(4))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[541]))
				F_BufFileWrite(m, v33, v8+int32(15), int32(1))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v41 = v39 - v40
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v41
					v44 = *(*int32)(unsafe.Add(mBase, _consts[541]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					F_BufFileWrite(m, v44, v40+v45, v41)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _consts[508]))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
						F_subxact_info_write(m, v51, l0)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[541]))
							F_BufFileClose(m, v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[541])) = int32(0)
								F_CommitTransactionCommand(m)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, _consts[539]))
									F_MemoryContextReset(m, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v19 = v11
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v2)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21 - v22 + int32(1)
		F_BufFileWrite(m, v19, v8+int32(8), int32(4))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _consts[541]))
			F_BufFileWrite(m, v33, v8+int32(15), int32(1))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
				v41 = v39 - v40
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v41
				v44 = *(*int32)(unsafe.Add(mBase, _consts[541]))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				F_BufFileWrite(m, v44, v40+v45, v41)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, _consts[508]))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
					F_subxact_info_write(m, v51, l0)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[541]))
						F_BufFileClose(m, v55)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[541])) = int32(0)
							F_CommitTransactionCommand(m)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, _consts[539]))
								F_MemoryContextReset(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_stream_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(365656)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(4508504)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(503810)
				F_errmsg(m, int32(319086), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(497876), int32(1444), int32(218077))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _consts[77])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
