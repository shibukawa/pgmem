package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IOContextForStrategy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = v8 - int32(1)
		if base.Ui32(int32(3)) <= base.Ui32(v10) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v29
				F_errmsg_internal(m, int32(_a_F_IOContextForStrategy_0), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_IOContextForStrategy_1), int32(824), int32(_a_F_IOContextForStrategy_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_c_F_IOContextForStrategy[0])))
			v18 = v15
			m.G0 = v6 + int32(16)
			return v18
		}
	} else {
		v18 = int32(3)
		m.G0 = v6 + int32(16)
		return v18
	}
}
func F_IdleSessionTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_IdleSessionTimeoutHandler[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_IdleSessionTimeoutHandler[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_IdleSessionTimeoutHandler[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	return
}
func F_InitSparseVector(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v7 = F_mul_size(m, int32(4), l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_add_size(m, int32(16), v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_mul_size(m, int32(4), l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_add_size(m, v11, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = F_palloc0(m, v16)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v16 << (uint(int32(2)) % 32)
						return v18
					}
				}
			}
		}
	}
}
func F_InitializeFastPathLocks(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v5 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeFastPathLocks[0]))
	if v8&(v8-v5) != 0 {
		v15 = v5 << (uint(int32(32)-base.I32_clz(v8)) % 32)
	} else {
		v15 = v8
	}
	if base.Ui32(v15) <= base.Ui32(int32(31)) {
		v18 = int32(31)
	} else {
		v18 = v15
	}
	if base.Ui32(int32(_a_F_InitializeFastPathLocks_0)) <= base.Ui32(v15) {
		v23 = int32(1024)
	} else {
		v23 = int32(base.Ui32(v18) >> (uint(int32(4)) % 32))
	}
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeFastPathLocks[1])) = v23
	return
}
func F_InitializeLatchWaitSet(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v5 = F_CreateWaitEventSet(m, int32(0), int32(2))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[0])) = v5
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[1]))
		F_AddWaitEventToSet(m, v5, int32(1), int32(-1), v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[2])))
			if v15 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[0]))
				F_AddWaitEventToSet(m, v19, int32(32), int32(-1), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
}
func F_IsAbortedTransactionBlockState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsAbortedTransactionBlockState[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return base.B2i32((v3-int32(7))&int32(-9) == int32(0))
}
func F__intbig_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F__intbig_out_0), int32(46), int32(_a_F__intbig_out_1), int32(_a_F__intbig_out_2), int32(_a_F__intbig_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_i4toi2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v2-int32(_a_F_i4toi2_0)) <= base.Ui32(int32(-65537)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_i4toi2_1), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_i4toi2_2), int32(384), int32(_a_F_i4toi2_3))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v2
	}
}
func F_iclikesel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13990(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_icu_language_tag(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_icu_language_tag_0), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_icu_language_tag_1), int32(1599), int32(_a_F_icu_language_tag_2))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_inc_lex_level(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = v5 + int32(1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v7
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v7 < v13 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v17 = v15 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v17
			v51 = v17
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v54+v51<<(uint(int32(2))%32)))) = int32(0)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v21 = v13 - int32(-64)
			v24 = F_repalloc(m, v19, v21*int32(10))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v24
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				v32 = F_repalloc(m, v29, v21<<(uint(int32(2))%32))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v32
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					v38 = F_repalloc(m, v37, v21)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v38
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v42))) = v21
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v45 = int32(1)
						v46 = v44 + v45
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v46
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v48 != v45 {
						} else {
							v51 = v46
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v54+v51<<(uint(int32(2))%32)))) = int32(0)
						}
						return
					}
				}
			}
		}
	}
}
func F_inetnot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = F_palloc0(m, int32(22))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v22 = v19 + v21
			v24 = v19 + int32(4)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			if v25&v21 != 0 {
				v28 = v22
			} else {
				v28 = v24
			}
			v30 = v28 + int32(2)
			v31 = int32(1)
			v32 = v14 + v31
			v34 = v14 + int32(4)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v37 = v35 & v31
			if v37 != 0 {
				v38 = v32
			} else {
				v38 = v34
			}
			v40 = v38 + int32(2)
			if v37 != 0 {
				v45 = int32(1)
			} else {
				v45 = int32(4)
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v45))))
			if v47 == int32(2) {
				v50 = int32(3)
			} else {
				v50 = int32(15)
			}
			v51 = v50
			for {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v40))))
				v66 = int32(-1)
				v67 = v65 ^ v66
				*(*uint8)(unsafe.Add(mBase, uint32(v51+v30))) = uint8(v67)
				v70 = v51 - int32(1)
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v40))))
				v75 = v73 ^ v66
				*(*uint8)(unsafe.Add(mBase, uint32(v30+v70))) = uint8(v75)
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v38))))
				v81 = v79 ^ v66
				*(*uint8)(unsafe.Add(mBase, uint32(v51+v28))) = uint8(v81)
				v84 = v51 - int32(3)
				v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v40))))
				v89 = v87 ^ v66
				*(*uint8)(unsafe.Add(mBase, uint32(v30+v84))) = uint8(v89)
				if v84 != 0 {
					v51 = v51 - int32(4)
					continue
				} else {
					break
				}
				break
			}
			v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			if v93&int32(1) != 0 {
				v96 = v22
			} else {
				v96 = v24
			}
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v97&int32(1) != 0 {
				v100 = v32
			} else {
				v100 = v34
			}
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)) = uint8(v101)
			v103 = int32(1)
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			if v105&v103 != 0 {
				v108 = v103
			} else {
				v108 = int32(4)
			}
			v110 = int32(1)
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v112&v110 != 0 {
				v115 = v110
			} else {
				v115 = int32(4)
			}
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v115))))
			*(*uint8)(unsafe.Add(mBase, uint32(v19+v108))) = uint8(v117)
			if v117 == int32(2) {
				v123 = int32(40)
			} else {
				v123 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v123
			return v19
		}
	}
}
func F_init_dummy_sjinfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+28)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+43)) = int32(0)
	return
}
func F_init_params(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int64
	_ = v426
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v433 int64
	_ = v433
	var v435 int32
	_ = v435
	var v436 int64
	_ = v436
	var v440 int64
	_ = v440
	var v443 int64
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v499 int32
	_ = v499
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int64
	_ = v596
	var v597 int32
	_ = v597
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int64
	_ = v647
	var v650 int32
	_ = v650
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int64
	_ = v710
	var v715 int64
	_ = v715
	var v720 int32
	_ = v720
	var v723 int64
	_ = v723
	var v724 int32
	_ = v724
	var v734 int64
	_ = v734
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v742 int64
	_ = v742
	var v743 int64
	_ = v743
	var v749 int64
	_ = v749
	var v750 int32
	_ = v750
	var v761 int64
	_ = v761
	var v763 int64
	_ = v763
	var v764 int32
	_ = v764
	var v766 int64
	_ = v766
	var v769 int64
	_ = v769
	var v770 int64
	_ = v770
	var v775 int64
	_ = v775
	var v776 int64
	_ = v776
	var v778 int64
	_ = v778
	var v780 int64
	_ = v780
	var v782 int32
	_ = v782
	var v783 int64
	_ = v783
	var v784 int32
	_ = v784
	var v785 int64
	_ = v785
	var v786 int32
	_ = v786
	var v793 int64
	_ = v793
	var v794 int32
	_ = v794
	var v797 int64
	_ = v797
	var v798 int64
	_ = v798
	var v800 int64
	_ = v800
	var v802 int64
	_ = v802
	var v803 int32
	_ = v803
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int64
	_ = v824
	var v825 int64
	_ = v825
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int64
	_ = v845
	var v846 int64
	_ = v846
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int64
	_ = v866
	var v867 int64
	_ = v867
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int64
	_ = v885
	var v886 int64
	_ = v886
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int64
	_ = v906
	var v907 int64
	_ = v907
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int64
	_ = v927
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int64
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int64
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	v11 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(144)
	m.G0 = v25
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v11
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errorConflictingDefElem(m, v60, l0)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L116
	} else {
		goto L300
	}
L2:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v707 - int32(21) {
	case 0:
		goto L193
	default:
		goto L192
	case 2:
		goto L194
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	v697 = v673
	v702 = v678
	v704 = v680
	v705 = v681
	v706 = v682
	goto L2
L4:
	;
	if v634 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L5:
	;
	if l3 != 0 {
		v634 = v610
		v635 = v611
		v640 = v616
		v642 = v618
		v643 = v619
		v644 = v620
		goto L4
	} else {
		goto L181
	}
L6:
	;
	if v587 == int32(0) {
		v610 = v580
		v611 = v581
		v616 = v586
		v618 = v588
		v619 = v589
		v620 = v590
		goto L5
	} else {
		goto L178
	}
L7:
	;
	v567 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+48)) = uint8(v567)
	v580 = v556
	v581 = v557
	v586 = v562
	v587 = v563
	v588 = v564
	v589 = v565
	v590 = v566
	goto L6
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v31 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	goto L10
L10:
	;
	if l3 == int32(0) {
		v610 = v11
		v611 = v11
		v616 = v11
		v618 = v11
		v619 = v11
		v620 = v11
		goto L5
	} else {
		goto L177
	}
L11:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+48)) = uint8(v533)
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	v580 = v530
	v581 = v531
	v586 = v397
	v587 = v398
	v588 = v399
	v589 = v400
	v590 = v401
	goto L6
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(1)
	v526 = int32(0)
	if v395 != 0 {
		v530 = v526
		v531 = v526
		goto L11
	} else {
		goto L176
	}
L13:
	;
	v498 = F_defGetInt64(m, v394)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L116
	} else {
		goto L166
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L116
	} else {
		goto L159
	}
L15:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L116
	} else {
		goto L155
	}
L16:
	;
	v46 = v11
	v48 = v11
	v49 = v11
	v50 = v11
	v51 = v11
	v52 = v11
	v53 = v11
	v54 = v11
	v55 = v11
	goto L19
L17:
	;
	v394 = v11
	v395 = v11
	v396 = v11
	v397 = v11
	v398 = v11
	v399 = v11
	v400 = v11
	v401 = v11
	goto L18
L18:
	;
	if l3 != 0 {
		goto L134
	} else {
		goto L135
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v46<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v62 != int32(97) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v394 = v368
	v395 = v369
	v396 = v370
	v397 = v371
	v398 = v372
	v399 = v373
	v400 = v374
	v401 = v375
	goto L18
L21:
	;
	v377 = v46 + int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v377 < v378 {
		v46 = v377
		v48 = v368
		v49 = v369
		v50 = v370
		v51 = v371
		v52 = v372
		v53 = v373
		v54 = v374
		v55 = v375
		goto L19
	} else {
		goto L129
	}
L22:
	;
	v366 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v366)
	v368 = v48
	v369 = v49
	v370 = v60
	v371 = v51
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	goto L21
L23:
	;
	v71 = int32(_a_F_init_params_0)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[0])))
	if base.B2i32(v74 == int32(0))|base.B2i32(v74 != v77) != 0 {
		v95 = v74
		v96 = v77
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v65 != int32(115) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v68 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if v50 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	if v95-v96 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	v80 = v61
	v81 = v71
	goto L31
L31:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v85
		v96 = v84
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v95 = v85
	v96 = v84
	goto L29
L33:
	;
	v88 = int32(1)
	if v85 == v84 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v48 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v102 = int32(_a_F_init_params_1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[1])))
	if base.B2i32(v105 == int32(0))|base.B2i32(v105 != v108) != 0 {
		v126 = v105
		v127 = v108
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v100)
	v368 = v60
	v369 = v49
	v370 = v50
	v371 = v51
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	goto L21
L39:
	;
	if v126-v127 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	v111 = v61
	v112 = v102
	goto L42
L42:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v116 == int32(0) {
		v126 = v116
		v127 = v115
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v126 = v116
	v127 = v115
	goto L40
L44:
	;
	v119 = int32(1)
	if v116 == v115 {
		v111 = v111 + v119
		v112 = v112 + v119
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v54 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v133 = int32(_a_F_init_params_2)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[2])))
	if base.B2i32(v136 == int32(0))|base.B2i32(v136 != v139) != 0 {
		v157 = v136
		v158 = v139
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v131)
	v368 = v48
	v369 = v49
	v370 = v50
	v371 = v51
	v372 = v52
	v373 = v53
	v374 = v60
	v375 = v55
	goto L21
L50:
	;
	if v157-v158 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v142 = v61
	v143 = v133
	goto L53
L53:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v147 == int32(0) {
		v157 = v147
		v158 = v146
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v157 = v147
	v158 = v146
	goto L51
L55:
	;
	v150 = int32(1)
	if v147 == v146 {
		v142 = v142 + v150
		v143 = v143 + v150
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if v51 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v164 = int32(_a_F_init_params_3)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[3])))
	if base.B2i32(v167 == int32(0))|base.B2i32(v167 != v170) != 0 {
		v188 = v167
		v189 = v170
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v162)
	v368 = v48
	v369 = v49
	v370 = v50
	v371 = v60
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	goto L21
L61:
	;
	if v188-v189 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	v173 = v61
	v174 = v164
	goto L64
L64:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	if v178 == int32(0) {
		v188 = v178
		v189 = v177
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v188 = v178
	v189 = v177
	goto L62
L66:
	;
	v181 = int32(1)
	if v178 == v177 {
		v173 = v173 + v181
		v174 = v174 + v181
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v52 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v195 = int32(_a_F_init_params_4)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[4])))
	if base.B2i32(v198 == int32(0))|base.B2i32(v198 != v201) != 0 {
		v219 = v198
		v220 = v201
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v193)
	v368 = v48
	v369 = v49
	v370 = v50
	v371 = v51
	v372 = v60
	v373 = v53
	v374 = v54
	v375 = v55
	goto L21
L72:
	;
	if v219-v220 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	goto L72
L74:
	;
	v204 = v61
	v205 = v195
	goto L75
L75:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	if v209 == int32(0) {
		v219 = v209
		v220 = v208
		goto L73
	} else {
		goto L77
	}
L76:
	;
	v219 = v209
	v220 = v208
	goto L73
L77:
	;
	v212 = int32(1)
	if v209 == v208 {
		v204 = v204 + v212
		v205 = v205 + v212
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if v53 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v226 = int32(_a_F_init_params_5)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[5])))
	if base.B2i32(v229 == int32(0))|base.B2i32(v229 != v232) != 0 {
		v250 = v229
		v251 = v232
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v224)
	v368 = v48
	v369 = v49
	v370 = v50
	v371 = v51
	v372 = v52
	v373 = v60
	v374 = v54
	v375 = v55
	goto L21
L83:
	;
	if v250-v251 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	v235 = v61
	v236 = v226
	goto L86
L86:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	if v240 == int32(0) {
		v250 = v240
		v251 = v239
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v250 = v240
	v251 = v239
	goto L84
L88:
	;
	v243 = int32(1)
	if v240 == v239 {
		v235 = v235 + v243
		v236 = v236 + v243
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	if v55 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v257 = int32(_a_F_init_params_6)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[6])))
	if base.B2i32(v260 == int32(0))|base.B2i32(v260 != v263) != 0 {
		v281 = v260
		v282 = v263
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v255)
	v368 = v48
	v369 = v49
	v370 = v50
	v371 = v51
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v60
	goto L21
L94:
	;
	if v281-v282 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	goto L94
L96:
	;
	v266 = v61
	v267 = v257
	goto L97
L97:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	if v271 == int32(0) {
		v281 = v271
		v282 = v270
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v281 = v271
	v282 = v270
	goto L95
L99:
	;
	v274 = int32(1)
	if v271 == v270 {
		v266 = v266 + v274
		v267 = v267 + v274
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	if v49 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v288 = int32(_a_F_init_params_7)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[7])))
	if base.B2i32(v291 == int32(0))|base.B2i32(v291 != v294) != 0 {
		v312 = v291
		v313 = v294
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v286)
	v368 = v48
	v369 = v60
	v370 = v50
	v371 = v51
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	goto L21
L105:
	;
	if v312-v313 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	v297 = v61
	v298 = v288
	goto L108
L108:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v302 == int32(0) {
		v312 = v302
		v313 = v301
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v312 = v302
	v313 = v301
	goto L106
L110:
	;
	v305 = int32(1)
	if v302 == v301 {
		v297 = v297 + v305
		v298 = v298 + v305
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	if v317 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v321 = int32(_a_F_init_params_8)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[8])))
	if base.B2i32(v324 == int32(0))|base.B2i32(v324 != v327) != 0 {
		v345 = v324
		v346 = v327
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v318 = F_defGetQualifiedName(m, v60)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	return
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v318
	v368 = v48
	v369 = v49
	v370 = v50
	v371 = v51
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	goto L21
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L116
	} else {
		goto L125
	}
L119:
	;
	goto L118
L120:
	;
	v330 = v61
	v331 = v321
	goto L121
L121:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+1)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v335 == int32(0) {
		v345 = v335
		v346 = v334
		goto L119
	} else {
		goto L123
	}
L122:
	;
	v345 = v335
	v346 = v334
	goto L119
L123:
	;
	v338 = int32(1)
	if v335 == v334 {
		v330 = v330 + v338
		v331 = v331 + v338
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v345-v346 == int32(0) {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+128)) = v354
	F_errmsg_internal(m, int32(_a_F_init_params_9), v25+int32(128))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L116
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1362), int32(_a_F_init_params_11))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L116
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	goto L20
L130:
	;
	if v394 != 0 {
		v496 = v458
		v497 = v459
		goto L13
	} else {
		goto L153
	}
L131:
	;
	v455 = int32(0)
	v458 = v455
	v459 = v455
	goto L130
L132:
	;
	v453 = int32(0)
	v496 = v453
	v497 = v453
	goto L13
L133:
	;
	v408 = F_defGetTypeName(m, v396)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L116
	} else {
		goto L140
	}
L134:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	if v396 != 0 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v396 == int32(0) {
		goto L131
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(20)
	if v394 != 0 {
		goto L132
	} else {
		goto L138
	}
L138:
	;
	goto L12
L139:
	;
	goto L133
L140:
	;
	v410 = F_typenameTypeId(m, l0, v408)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L116
	} else {
		goto L141
	}
L141:
	;
	if base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v410-int32(20)))&base.B2i32(v410 != int32(23)) != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	if l3 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v421 = int32(0)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v423 - int32(20) {
	case 0:
		goto L147
	case 1:
		goto L149
	default:
		v446 = v421
		v447 = v421
		goto L146
	case 3:
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v410
	if v394 == int32(0) {
		goto L12
	} else {
		goto L152
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v410
	v458 = v446
	v459 = v447
	goto L130
L147:
	;
	v440 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v443 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v446 = base.B2i32(v443 == int64(9223372036854775807))
	v447 = base.B2i32(v440 == int64(-9223372036854775807-1))
	goto L146
L148:
	;
	v433 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v435 = base.B2i32(v433 == int64(2147483647))
	v436 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v436 != int64(-2147483648) {
		v446 = v435
		v447 = v421
		goto L146
	} else {
		goto L151
	}
L149:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v428 = base.B2i32(v426 == int64(32767))
	v429 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v429 != int64(-32768) {
		v446 = v428
		v447 = v421
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v446 = v428
	v447 = int32(1)
	goto L146
L151:
	;
	v446 = v435
	v447 = int32(1)
	goto L146
L152:
	;
	goto L132
L153:
	;
	if v395 == int32(0) {
		v580 = v458
		v581 = v459
		v586 = v397
		v587 = v398
		v588 = v399
		v589 = v400
		v590 = v401
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v530 = v458
	v531 = v459
	goto L11
L155:
	;
	F_errmsg(m, int32(_a_F_init_params_12), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L116
	} else {
		goto L156
	}
L156:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	F_parser_errposition(m, l0, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L116
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1358), int32(_a_F_init_params_11))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L116
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L116
	} else {
		goto L160
	}
L160:
	;
	if l2 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v486 = int32(_a_F_init_params_13)
	goto L163
L162:
	;
	v486 = int32(_a_F_init_params_14)
	goto L163
L163:
	;
	F_errmsg(m, v486, int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L116
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1384), int32(_a_F_init_params_11))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L116
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = v498
	if v498 != int64(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	if v395 != 0 {
		v530 = v496
		v531 = v497
		goto L11
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L116
	} else {
		goto L172
	}
L170:
	;
	if l3 == int32(0) {
		v580 = v496
		v581 = v497
		v586 = v397
		v587 = v398
		v588 = v399
		v589 = v400
		v590 = v401
		goto L6
	} else {
		goto L171
	}
L171:
	;
	v556 = v496
	v557 = v497
	v562 = v397
	v563 = v398
	v564 = v399
	v565 = v400
	v566 = v401
	goto L7
L172:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L116
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(_a_F_init_params_15), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L116
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1418), int32(_a_F_init_params_11))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L116
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v556 = v526
	v557 = v526
	v562 = v397
	v563 = v398
	v564 = v399
	v565 = v400
	v566 = v401
	goto L7
L177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(20)
	v556 = v11
	v557 = v11
	v562 = v11
	v563 = v11
	v564 = v11
	v565 = v11
	v566 = v11
	goto L7
L178:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	if v593 == int32(0) {
		v634 = v580
		v635 = v581
		v640 = v586
		v642 = v588
		v643 = v589
		v644 = v590
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v596 = F_defGetInt64(m, v587)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L116
	} else {
		goto L180
	}
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = v596
	v673 = v581
	v678 = v586
	v680 = v588
	v681 = v589
	v682 = v590
	goto L3
L181:
	;
	if v610 == int32(0) {
		v697 = v611
		v702 = v616
		v704 = v618
		v705 = v619
		v706 = v620
		goto L2
	} else {
		goto L182
	}
L182:
	;
	v634 = v610
	v635 = v611
	v640 = v616
	v642 = v618
	v643 = v619
	v644 = v620
	goto L4
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(-1)
	v673 = v635
	v678 = v640
	v680 = v642
	v681 = v643
	v682 = v644
	goto L3
L184:
	;
	v647 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if v647 <= int64(0) {
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v650 - int32(21) {
	case 0:
		goto L190
	default:
		goto L188
	case 2:
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(9223372036854775807)
	v673 = v635
	v678 = v640
	v680 = v642
	v681 = v643
	v682 = v644
	goto L3
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(2147483647)
	v673 = v635
	v678 = v640
	v680 = v642
	v681 = v643
	v682 = v644
	goto L3
L190:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(32767)
	v673 = v635
	v678 = v640
	v680 = v642
	v681 = v643
	v682 = v644
	goto L3
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L116
	} else {
		goto L295
	}
L192:
	;
	if v704 != 0 {
		goto L200
	} else {
		goto L201
	}
L193:
	;
	v715 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if base.Ui64(v715-int64(32768)) < base.Ui64(int64(-65536)) {
		goto L191
	} else {
		goto L196
	}
L194:
	;
	v710 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if base.Ui64(int64(-4294967297)) < base.Ui64(v710-int64(2147483648)) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L191
L196:
	;
	goto L192
L197:
	;
	v749 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v750 - int32(21) {
	case 0:
		goto L220
	default:
		goto L219
	case 2:
		goto L221
	}
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	goto L197
L199:
	;
	if v707 == int32(23) {
		goto L206
	} else {
		goto L207
	}
L200:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v704)+12))
	if v720 == int32(0) {
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if l3|v697 != int32(1) {
		goto L197
	} else {
		goto L205
	}
L203:
	;
	v723 = F_defGetInt64(m, v704)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L116
	} else {
		goto L204
	}
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = v723
	goto L198
L205:
	;
	goto L199
L206:
	;
	v734 = int64(-2147483648)
	goto L208
L207:
	;
	v734 = int64(-9223372036854775807 - 1)
	goto L208
L208:
	;
	if v707 == int32(21) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v737 = int64(-32768)
	goto L211
L210:
	;
	v737 = v734
	goto L211
L211:
	;
	v739 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if int64(0) <= v739 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v742 = int64(1)
	goto L214
L213:
	;
	v742 = v737
	goto L214
L214:
	;
	if v697 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v743 = v737
	goto L217
L216:
	;
	v743 = v742
	goto L217
L217:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = v743
	goto L198
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L116
	} else {
		goto L290
	}
L219:
	;
	v761 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v749 < v761 {
		goto L229
	} else {
		goto L230
	}
L220:
	;
	if base.Ui64(v749-int64(32768)) < base.Ui64(int64(-65536)) {
		goto L218
	} else {
		goto L223
	}
L221:
	;
	if base.Ui64(int64(-4294967297)) < base.Ui64(v749-int64(2147483648)) {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L218
L223:
	;
	goto L219
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L116
	} else {
		goto L286
	}
L225:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L116
	} else {
		goto L282
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L116
	} else {
		goto L278
	}
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L116
	} else {
		goto L274
	}
L228:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L116
	} else {
		goto L270
	}
L229:
	;
	if v705 != 0 {
		goto L234
	} else {
		goto L235
	}
L230:
	;
	goto L231
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L116
	} else {
		goto L266
	}
L232:
	;
	v780 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v780 < v778 {
		goto L227
	} else {
		goto L245
	}
L233:
	;
	if v775 < v776 {
		goto L228
	} else {
		goto L244
	}
L234:
	;
	v763 = F_defGetInt64(m, v705)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L116
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if l3 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v763
	v766 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v775 = v763
	v776 = v766
	goto L233
L238:
	;
	v769 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v775 = v769
	v776 = v749
	goto L233
L239:
	;
	goto L240
L240:
	;
	v770 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if int64(0) < v770 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v749
	v778 = v749
	goto L232
L242:
	;
	goto L243
L243:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v761
	v775 = v761
	v776 = v749
	goto L233
L244:
	;
	v778 = v775
	goto L232
L245:
	;
	if v702 != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v798 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v797 < v798 {
		goto L226
	} else {
		goto L257
	}
L247:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v702)+12))
	if v782 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L249
L249:
	;
	if l3 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L250:
	;
	v783 = F_defGetInt64(m, v702)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L116
	} else {
		goto L253
	}
L251:
	;
	v785 = v778
	goto L252
L252:
	;
	v786 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)) = uint8(v786)
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v785
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	v797 = v785
	goto L246
L253:
	;
	v785 = v783
	goto L252
L254:
	;
	v793 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v797 = v793
	goto L246
L255:
	;
	goto L256
L256:
	;
	v794 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)) = uint8(v794)
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v778
	v797 = v778
	goto L246
L257:
	;
	v800 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v800 < v797 {
		goto L225
	} else {
		goto L258
	}
L258:
	;
	if v706 != 0 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	m.G0 = v25 + int32(144)
	return
L260:
	;
	v802 = F_defGetInt64(m, v706)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L116
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	if l3 == int32(0) {
		goto L259
	} else {
		goto L265
	}
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+40)) = v802
	if v802 <= int64(0) {
		goto L224
	} else {
		goto L264
	}
L264:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	goto L259
L265:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+40)) = int64(1)
	goto L259
L266:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L116
	} else {
		goto L267
	}
L267:
	;
	v824 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v825 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+88)) = v825
	*(*int64)(unsafe.Add(mBase, uint32(v25)+80)) = v824
	F_errmsg(m, int32(_a_F_init_params_16), v25+int32(80))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L116
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1508), int32(_a_F_init_params_11))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L116
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L116
	} else {
		goto L271
	}
L271:
	;
	v845 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v846 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+72)) = v846
	*(*int64)(unsafe.Add(mBase, uint32(v25)+64)) = v845
	F_errmsg(m, int32(_a_F_init_params_17), v25-int32(-64))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L116
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1529), int32(_a_F_init_params_11))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L116
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L116
	} else {
		goto L275
	}
L275:
	;
	v866 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v867 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v866
	F_errmsg(m, int32(_a_F_init_params_18), v25)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L116
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1535), int32(_a_F_init_params_11))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L116
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L116
	} else {
		goto L279
	}
L279:
	;
	v885 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v886 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v886
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v885
	F_errmsg(m, int32(_a_F_init_params_19), v25+int32(16))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L116
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1559), int32(_a_F_init_params_11))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L116
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L116
	} else {
		goto L283
	}
L283:
	;
	v906 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v907 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v907
	*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v906
	F_errmsg(m, int32(_a_F_init_params_20), v25+int32(32))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L116
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1565), int32(_a_F_init_params_11))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L116
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L116
	} else {
		goto L287
	}
L287:
	;
	v927 = *(*int64)(unsafe.Add(mBase, uint32(l4)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v927
	F_errmsg(m, int32(_a_F_init_params_21), v25+int32(48))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L116
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1575), int32(_a_F_init_params_11))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L116
	} else {
		goto L289
	}
L289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L290:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L116
	} else {
		goto L291
	}
L291:
	;
	v946 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v948 = F_format_type_be(m, v947)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L116
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = v948
	*(*int64)(unsafe.Add(mBase, uint32(v25)+96)) = v946
	F_errmsg(m, int32(_a_F_init_params_22), v25+int32(96))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L116
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1500), int32(_a_F_init_params_11))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L116
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L116
	} else {
		goto L296
	}
L296:
	;
	v969 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v971 = F_format_type_be(m, v970)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L116
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+120)) = v971
	*(*int64)(unsafe.Add(mBase, uint32(v25)+112)) = v969
	F_errmsg(m, int32(_a_F_init_params_23), v25+int32(112))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L116
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1468), int32(_a_F_init_params_11))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L116
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_init_ps_display(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	if l0 == int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_init_ps_display[0]))
		if base.Ui32(v5) <= base.Ui32(int32(17)) {
		} else {
		}
	} else {
	}
	return
}
func F_init_span(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	v6 = l5
	if l1 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1468))
		if v13 != v15 {
			v20 = F_LWLockAcquire(m, v14+int32(1476), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v24+int32(1476))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v32 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
						v35 = l0 + v32*int32(20)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
						if v36 != 0 {
							v40 = v36
							v47 = v40 + l1&int32(134217727)
							v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							if v51 != 0 {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1468))
								if v52 != v54 {
									v59 = F_LWLockAcquire(m, v53+int32(1476), int32(0))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v63+int32(1476))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
												v74 = l0 + v71*int32(20)
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
												if v75 != 0 {
													v79 = v75
													*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
													v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
													v90 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
													*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
													*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
													switch v6 {
													case 0:
														v100 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
														v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
														v108 = v103 - v100
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
														v110 = v108
													case 1:
														v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
														v110 = v99
													default:
														v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
														v108 = v107
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
														v110 = v108
													}
													v111 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
													v114 = int32(_a_F_init_span_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
													return
												} else {
													v76 = F_get_segment_by_index(m, l0, v71)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return
													} else {
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
														v79 = v78
														*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
														v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
														v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
														v90 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
														*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
														*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
														switch v6 {
														case 0:
															v100 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
															v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
															v108 = v103 - v100
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
															v110 = v108
														case 1:
															v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
															v110 = v99
														default:
															v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
															v108 = v107
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
															v110 = v108
														}
														v111 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
														v114 = int32(_a_F_init_span_2)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
														return
													}
												}
											}
										}
									}
								} else {
									v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
									v74 = l0 + v71*int32(20)
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									if v75 != 0 {
										v79 = v75
										*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
										v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
										v90 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
										*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
										*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
										switch v6 {
										case 0:
											v100 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
											v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
											v108 = v103 - v100
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
											v110 = v108
										case 1:
											v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
											v110 = v99
										default:
											v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
											v108 = v107
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
											v110 = v108
										}
										v111 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
										v114 = int32(_a_F_init_span_2)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
										return
									} else {
										v76 = F_get_segment_by_index(m, l0, v71)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
											v79 = v78
											*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v90 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
											*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
											switch v6 {
											case 0:
												v100 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
												v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
												v108 = v103 - v100
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											case 1:
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
												v110 = v99
											default:
												v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
												v108 = v107
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											}
											v111 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
											v114 = int32(_a_F_init_span_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
											return
										}
									}
								}
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v90 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
								switch v6 {
								case 0:
									v100 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
									v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
									v108 = v103 - v100
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								case 1:
									v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
									v110 = v99
								default:
									v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
									v108 = v107
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								}
								v111 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
								v114 = int32(_a_F_init_span_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
								return
							}
						} else {
							v37 = F_get_segment_by_index(m, l0, v32)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
								v40 = v39
								v47 = v40 + l1&int32(134217727)
								v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								if v51 != 0 {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1468))
									if v52 != v54 {
										v59 = F_LWLockAcquire(m, v53+int32(1476), int32(0))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v63+int32(1476))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return
												} else {
													v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
													v74 = l0 + v71*int32(20)
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
													if v75 != 0 {
														v79 = v75
														*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
														v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
														v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
														v90 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
														*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
														*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
														switch v6 {
														case 0:
															v100 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
															v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
															v108 = v103 - v100
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
															v110 = v108
														case 1:
															v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
															v110 = v99
														default:
															v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
															v108 = v107
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
															v110 = v108
														}
														v111 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
														v114 = int32(_a_F_init_span_2)
														*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
														return
													} else {
														v76 = F_get_segment_by_index(m, l0, v71)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
															v79 = v78
															*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
															v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
															v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
															v90 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
															*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
															*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
															*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
															*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
															switch v6 {
															case 0:
																v100 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
																v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
																v108 = v103 - v100
																*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
																v110 = v108
															case 1:
																v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
																v110 = v99
															default:
																v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
																v108 = v107
																*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
																v110 = v108
															}
															v111 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
															v114 = int32(_a_F_init_span_2)
															*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
															return
														}
													}
												}
											}
										}
									} else {
										v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
										v74 = l0 + v71*int32(20)
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										if v75 != 0 {
											v79 = v75
											*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v90 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
											*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
											switch v6 {
											case 0:
												v100 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
												v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
												v108 = v103 - v100
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											case 1:
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
												v110 = v99
											default:
												v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
												v108 = v107
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											}
											v111 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
											v114 = int32(_a_F_init_span_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
											return
										} else {
											v76 = F_get_segment_by_index(m, l0, v71)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
												v79 = v78
												*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
												v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
												v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												v90 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
												*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
												*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
												*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
												switch v6 {
												case 0:
													v100 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
													v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
													v108 = v103 - v100
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
													v110 = v108
												case 1:
													v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
													v110 = v99
												default:
													v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
													v108 = v107
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
													v110 = v108
												}
												v111 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
												v114 = int32(_a_F_init_span_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
												return
											}
										}
									}
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v90 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
									switch v6 {
									case 0:
										v100 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
										v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
										v108 = v103 - v100
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
										v110 = v108
									case 1:
										v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
										v110 = v99
									default:
										v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
										v108 = v107
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
										v110 = v108
									}
									v111 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
									v114 = int32(_a_F_init_span_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v32 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
			v35 = l0 + v32*int32(20)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
			if v36 != 0 {
				v40 = v36
				v47 = v40 + l1&int32(134217727)
				v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v51 != 0 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1468))
					if v52 != v54 {
						v59 = F_LWLockAcquire(m, v53+int32(1476), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v63+int32(1476))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
									v74 = l0 + v71*int32(20)
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									if v75 != 0 {
										v79 = v75
										*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
										v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
										v90 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
										*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
										*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
										switch v6 {
										case 0:
											v100 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
											v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
											v108 = v103 - v100
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
											v110 = v108
										case 1:
											v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
											v110 = v99
										default:
											v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
											v108 = v107
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
											v110 = v108
										}
										v111 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
										v114 = int32(_a_F_init_span_2)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
										return
									} else {
										v76 = F_get_segment_by_index(m, l0, v71)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
											v79 = v78
											*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v90 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
											*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
											switch v6 {
											case 0:
												v100 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
												v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
												v108 = v103 - v100
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											case 1:
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
												v110 = v99
											default:
												v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
												v108 = v107
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											}
											v111 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
											v114 = int32(_a_F_init_span_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
											return
										}
									}
								}
							}
						}
					} else {
						v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
						v74 = l0 + v71*int32(20)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
						if v75 != 0 {
							v79 = v75
							*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							v90 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
							*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
							switch v6 {
							case 0:
								v100 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
								v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
								v108 = v103 - v100
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
								v110 = v108
							case 1:
								v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
								v110 = v99
							default:
								v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
								v108 = v107
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
								v110 = v108
							}
							v111 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
							v114 = int32(_a_F_init_span_2)
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
							return
						} else {
							v76 = F_get_segment_by_index(m, l0, v71)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								v79 = v78
								*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v90 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
								switch v6 {
								case 0:
									v100 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
									v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
									v108 = v103 - v100
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								case 1:
									v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
									v110 = v99
								default:
									v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
									v108 = v107
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								}
								v111 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
								v114 = int32(_a_F_init_span_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
								return
							}
						}
					}
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v90 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
					*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
					switch v6 {
					case 0:
						v100 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
						v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
						v108 = v103 - v100
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
						v110 = v108
					case 1:
						v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
						v110 = v99
					default:
						v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
						v108 = v107
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
						v110 = v108
					}
					v111 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
					v114 = int32(_a_F_init_span_2)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
					return
				}
			} else {
				v37 = F_get_segment_by_index(m, l0, v32)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
					v40 = v39
					v47 = v40 + l1&int32(134217727)
					v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					if v51 != 0 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1468))
						if v52 != v54 {
							v59 = F_LWLockAcquire(m, v53+int32(1476), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v63+int32(1476))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
										v74 = l0 + v71*int32(20)
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										if v75 != 0 {
											v79 = v75
											*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v90 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
											*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
											switch v6 {
											case 0:
												v100 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
												v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
												v108 = v103 - v100
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											case 1:
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
												v110 = v99
											default:
												v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
												v108 = v107
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
												v110 = v108
											}
											v111 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
											v114 = int32(_a_F_init_span_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
											return
										} else {
											v76 = F_get_segment_by_index(m, l0, v71)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
												v79 = v78
												*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
												v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
												v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												v90 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
												*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
												*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
												*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
												switch v6 {
												case 0:
													v100 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
													v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
													v108 = v103 - v100
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
													v110 = v108
												case 1:
													v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
													v110 = v99
												default:
													v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
													v108 = v107
													*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
													v110 = v108
												}
												v111 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
												v114 = int32(_a_F_init_span_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
												return
											}
										}
									}
								}
							}
						} else {
							v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
							v74 = l0 + v71*int32(20)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
							if v75 != 0 {
								v79 = v75
								*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v90 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
								switch v6 {
								case 0:
									v100 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
									v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
									v108 = v103 - v100
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								case 1:
									v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
									v110 = v99
								default:
									v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
									v108 = v107
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								}
								v111 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
								v114 = int32(_a_F_init_span_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
								return
							} else {
								v76 = F_get_segment_by_index(m, l0, v71)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									v79 = v78
									*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v90 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
									switch v6 {
									case 0:
										v100 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
										v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
										v108 = v103 - v100
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
										v110 = v108
									case 1:
										v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
										v110 = v99
									default:
										v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
										v108 = v107
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
										v110 = v108
									}
									v111 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
									v114 = int32(_a_F_init_span_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
									return
								}
							}
						}
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						v90 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
						*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
						switch v6 {
						case 0:
							v100 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
							v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
							v108 = v103 - v100
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
							v110 = v108
						case 1:
							v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
							v110 = v99
						default:
							v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
							v108 = v107
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
							v110 = v108
						}
						v111 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
						v114 = int32(_a_F_init_span_2)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
						return
					}
				}
			}
		}
	} else {
		v47 = int32(0)
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		if v51 != 0 {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1468))
			if v52 != v54 {
				v59 = F_LWLockAcquire(m, v53+int32(1476), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v63+int32(1476))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
							v74 = l0 + v71*int32(20)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
							if v75 != 0 {
								v79 = v75
								*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v90 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
								switch v6 {
								case 0:
									v100 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
									v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
									v108 = v103 - v100
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								case 1:
									v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
									v110 = v99
								default:
									v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
									v108 = v107
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
									v110 = v108
								}
								v111 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
								v114 = int32(_a_F_init_span_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
								return
							} else {
								v76 = F_get_segment_by_index(m, l0, v71)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									v79 = v78
									*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v90 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
									switch v6 {
									case 0:
										v100 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
										v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
										v108 = v103 - v100
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
										v110 = v108
									case 1:
										v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
										v110 = v99
									default:
										v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
										v108 = v107
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
										v110 = v108
									}
									v111 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
									v114 = int32(_a_F_init_span_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
									return
								}
							}
						}
					}
				}
			} else {
				v71 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
				v74 = l0 + v71*int32(20)
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
				if v75 != 0 {
					v79 = v75
					*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v90 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
					*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
					switch v6 {
					case 0:
						v100 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
						v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
						v108 = v103 - v100
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
						v110 = v108
					case 1:
						v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
						v110 = v99
					default:
						v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
						v108 = v107
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
						v110 = v108
					}
					v111 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
					v114 = int32(_a_F_init_span_2)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
					return
				} else {
					v76 = F_get_segment_by_index(m, l0, v71)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
						v79 = v78
						*(*int32)(unsafe.Add(mBase, uint32(v79+v51&int32(134217727))+4)) = l1
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						v90 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
						*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
						switch v6 {
						case 0:
							v100 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
							v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
							v108 = v103 - v100
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
							v110 = v108
						case 1:
							v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
							v110 = v99
						default:
							v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
							v108 = v107
							*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
							v110 = v108
						}
						v111 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
						v114 = int32(_a_F_init_span_2)
						*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
						return
					}
				}
			}
		} else {
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2 - v86
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v90 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v90
			*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v89
			*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v90)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+20)) = uint16(v6)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = l3
			switch v6 {
			case 0:
				v100 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+22)) = uint16(v100)
				v103 = base.I32_div_u_s(int32(_a_F_init_span_0), v50)
				v108 = v103 - v100
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
				v110 = v108
			case 1:
				v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
				v110 = v99
			default:
				v107 = base.I32_div_u_s(int32(_a_F_init_span_1), v50)
				v108 = v107
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v108)
				v110 = v108
			}
			v111 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+30)) = uint16(v111)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+28)) = uint16(v110)
			v114 = int32(_a_F_init_span_2)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+26)) = uint16(v114)
			return
		}
	}
}
func F_initcap_wbnext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v6) <= base.Ui32(v5) {
		v234 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v234
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = v8 + v5
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v10 == int32(0) {
		v234 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = v10
	v16 = v9
	goto L4
L4:
	;
	v18 = v14 & int32(255)
	if int32(0) <= base.I32_extend8_s(v14) {
		v76 = v18
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v234 = v227
	goto L1
L6:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(128)) <= base.Ui32(v76) {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	if v18&int32(224) == int32(192) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v16))))
	v76 = v71&int32(63) | v68
	goto L6
L9:
	;
	v68 = v18 << (uint(int32(6)) % 32) & int32(1984)
	v69 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v18&int32(240) == int32(224) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v68 = v18<<(uint(int32(12))%32)&int32(_a_F_initcap_wbnext_0) | v39&int32(63)<<(uint(int32(6))%32)
	v69 = int32(2)
	goto L8
L13:
	;
	goto L14
L14:
	;
	if v18&int32(248) != int32(240) {
		v76 = int32(-1)
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v56 = int32(63)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v68 = v18<<(uint(int32(18))%32)&int32(_a_F_initcap_wbnext_1) | v55&v56<<(uint(int32(12))%32) | v61&v56<<(uint(int32(6))%32)
	v69 = int32(3)
	goto L8
L16:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v188 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L17:
	;
	v187 = v180
	goto L16
L18:
	;
	v180 = base.B2i32(v169&int32(255) == int32(9))
	goto L17
L19:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+uint32(_c_F_initcap_wbnext[0]))))
	v169 = v162
	goto L18
L20:
	;
	v187 = base.B2i32(base.Ui32(v76-int32(48)) < base.Ui32(int32(10)))
	goto L16
L21:
	;
	v87 = int32(1178)
	v88 = int32(0)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v142 = int32(1)
	v144 = v76 << (uint(v142) % 32)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+uint32(_c_F_initcap_wbnext[1]))))
	if v145&v142 != 0 {
		v180 = v142
		goto L17
	} else {
		goto L44
	}
L24:
	;
	v93 = base.I32_div_s(v87+v88, int32(2))
	v95 = v93 << (uint(int32(3)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+uint32(_c_F_initcap_wbnext[2])))
	if base.Ui32(v98) < base.Ui32(v76) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v77 != 0 {
		goto L20
	} else {
		goto L34
	}
L26:
	;
	if v110 <= v109 {
		v87 = v109
		v88 = v110
		goto L24
	} else {
		goto L33
	}
L27:
	;
	v109 = v87
	v110 = v93 + int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+uint32(_c_F_initcap_wbnext[3])))
	if base.Ui32(v104) <= base.Ui32(v76) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v187 = int32(1)
	goto L16
L31:
	;
	goto L32
L32:
	;
	v109 = v93 - int32(1)
	v110 = v88
	goto L26
L33:
	;
	goto L25
L34:
	;
	v116 = int32(3367)
	v117 = int32(0)
	goto L36
L35:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_initcap_wbnext[4]))))
	v169 = v141
	goto L18
L36:
	;
	v122 = base.I32_div_s(v116+v117, int32(2))
	v124 = v122 * int32(12)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_initcap_wbnext[5])))
	if base.Ui32(v127) < base.Ui32(v76) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v169 = int32(0)
	goto L18
L38:
	;
	if v138 <= v137 {
		v116 = v137
		v117 = v138
		goto L36
	} else {
		goto L43
	}
L39:
	;
	v137 = v116
	v138 = v122 + int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_initcap_wbnext[6])))
	if base.Ui32(v133) <= base.Ui32(v76) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	v137 = v122 - int32(1)
	v138 = v117
	goto L38
L43:
	;
	goto L37
L44:
	;
	if v77 == int32(0) {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	goto L20
L46:
	;
	if base.Ui32(v76) < base.Ui32(int32(128)) {
		v223 = int32(1)
		goto L57
	} else {
		goto L58
	}
L47:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v191 == v187 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v187)
	if base.Ui32(v76) < base.Ui32(int32(128)) {
		v208 = v193
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v208 + v195
	return v195
L52:
	;
	if base.Ui32(v76) < base.Ui32(int32(2048)) {
		v208 = int32(2)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v76) < base.Ui32(int32(_a_F_initcap_wbnext_2)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v207 = int32(3)
	goto L56
L55:
	;
	v207 = int32(4)
	goto L56
L56:
	;
	v208 = v207
	goto L51
L57:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v225 = v223 + v224
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v227) <= base.Ui32(v225) {
		v234 = v227
		goto L1
	} else {
		goto L63
	}
L58:
	;
	if base.Ui32(v76) < base.Ui32(int32(2048)) {
		v223 = int32(2)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v76) < base.Ui32(int32(_a_F_initcap_wbnext_2)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v222 = int32(3)
	goto L62
L61:
	;
	v222 = int32(4)
	goto L62
L62:
	;
	v223 = v222
	goto L57
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v230 = v229 + v225
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v231 != 0 {
		v14 = v231
		v16 = v230
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L5
}
func F_initial_cost_mergejoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v29 float64
	_ = v29
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 float64
	_ = v427
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v441 float64
	_ = v441
	var v443 float64
	_ = v443
	var v444 float64
	_ = v444
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 float64
	_ = v458
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 float64
	_ = v468
	var v469 int32
	_ = v469
	var v472 float64
	_ = v472
	var v474 float64
	_ = v474
	var v475 float64
	_ = v475
	var v486 int32
	_ = v486
	var v489 float64
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 float32
	_ = v493
	var v495 float64
	_ = v495
	var v497 float64
	_ = v497
	var v502 float64
	_ = v502
	var v507 float64
	_ = v507
	var v510 float64
	_ = v510
	var v511 float32
	_ = v511
	var v513 float64
	_ = v513
	var v515 float64
	_ = v515
	var v520 float64
	_ = v520
	var v525 float64
	_ = v525
	var v530 int32
	_ = v530
	var v533 float64
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 float32
	_ = v537
	var v539 float64
	_ = v539
	var v541 float64
	_ = v541
	var v546 float64
	_ = v546
	var v551 float64
	_ = v551
	var v554 float64
	_ = v554
	var v555 float32
	_ = v555
	var v557 float64
	_ = v557
	var v559 float64
	_ = v559
	var v564 float64
	_ = v564
	var v569 float64
	_ = v569
	var v574 float64
	_ = v574
	var v575 float64
	_ = v575
	var v581 float64
	_ = v581
	var v582 float64
	_ = v582
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 float64
	_ = v645
	var v647 float64
	_ = v647
	var v649 float64
	_ = v649
	var v651 float64
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v671 int32
	_ = v671
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 float64
	_ = v765
	var v768 int32
	_ = v768
	var v770 float64
	_ = v770
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v779 float64
	_ = v779
	var v782 int32
	_ = v782
	var v784 float64
	_ = v784
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
	var v790 float64
	_ = v790
	var v819 float64
	_ = v819
	var v820 float64
	_ = v820
	var v822 float64
	_ = v822
	var v825 float64
	_ = v825
	var v839 float64
	_ = v839
	var v844 float64
	_ = v844
	var v846 float64
	_ = v846
	var v847 float64
	_ = v847
	var v856 float64
	_ = v856
	var v860 float64
	_ = v860
	var v862 float64
	_ = v862
	var v863 float64
	_ = v863
	var v872 float64
	_ = v872
	var v876 float64
	_ = v876
	var v877 float64
	_ = v877
	var v878 float64
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v890 float64
	_ = v890
	var v891 float64
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 float64
	_ = v899
	var v900 int32
	_ = v900
	var v901 float64
	_ = v901
	var v902 float64
	_ = v902
	var v905 float64
	_ = v905
	var v907 float64
	_ = v907
	var v911 float64
	_ = v911
	var v912 float64
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v922 float64
	_ = v922
	var v924 int64
	_ = v924
	var v926 int64
	_ = v926
	var v927 float64
	_ = v927
	var v929 float64
	_ = v929
	var v935 float64
	_ = v935
	var v937 float64
	_ = v937
	var v940 int32
	_ = v940
	var v942 int64
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 float64
	_ = v950
	var v952 float64
	_ = v952
	var v953 float64
	_ = v953
	var v957 float64
	_ = v957
	var v960 float64
	_ = v960
	var v964 float64
	_ = v964
	var v971 float64
	_ = v971
	var v972 float64
	_ = v972
	var v974 float64
	_ = v974
	var v978 float64
	_ = v978
	var v982 float64
	_ = v982
	var v983 float64
	_ = v983
	var v986 int32
	_ = v986
	var v990 float64
	_ = v990
	var v993 int32
	_ = v993
	var v995 float64
	_ = v995
	var v1000 float64
	_ = v1000
	var v1001 float64
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 float64
	_ = v1007
	var v1008 float64
	_ = v1008
	var v1009 float64
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1020 float64
	_ = v1020
	var v1023 float64
	_ = v1023
	var v1024 float64
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 float64
	_ = v1026
	var v1029 float64
	_ = v1029
	var v1031 float64
	_ = v1031
	var v1035 float64
	_ = v1035
	var v1036 float64
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1046 float64
	_ = v1046
	var v1048 int64
	_ = v1048
	var v1050 int64
	_ = v1050
	var v1051 float64
	_ = v1051
	var v1053 float64
	_ = v1053
	var v1059 float64
	_ = v1059
	var v1061 float64
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int64
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 float64
	_ = v1074
	var v1076 float64
	_ = v1076
	var v1077 float64
	_ = v1077
	var v1081 float64
	_ = v1081
	var v1084 float64
	_ = v1084
	var v1088 float64
	_ = v1088
	var v1095 float64
	_ = v1095
	var v1096 float64
	_ = v1096
	var v1098 float64
	_ = v1098
	var v1102 float64
	_ = v1102
	var v1106 float64
	_ = v1106
	var v1107 float64
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 float64
	_ = v1114
	var v1117 float64
	_ = v1117
	var v1121 float64
	_ = v1121
	var v1122 float64
	_ = v1122
	var v1123 float64
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 float64
	_ = v1128
	var v1134 float64
	_ = v1134
	var v1142 float64
	_ = v1142
	var v1148 float64
	_ = v1148
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	v10 = int32(0)
	v29 = float64(0)
	v46 = m.G0
	v48 = v46 - int32(96)
	m.G0 = v48
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v52 = float64(1)
	if base.B2i32(l3 == v10)|base.B2i32(l2 == int32(2)) != 0 {
		v819 = v52
		v820 = v29
		v822 = v52
		v825 = v29
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L32
	} else {
		goto L246
	}
L2:
	;
	if base.F64_le(v50, float64(0)) != 0 {
		goto L180
	} else {
		goto L181
	}
L3:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v60 = l6
	goto L6
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v60 = v59
	goto L6
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if l7 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v65 = l7
	goto L9
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l5)+64))
	v65 = v64
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v63 != v68 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v71 != v73 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v75 != v76 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+16)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+16)))
	if v78 != v79 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+116))
	if v83 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+8))
	v709 = int32(0)
	if v706 == v709 {
		goto L148
	} else {
		goto L149
	}
L15:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v197 = m.G0
	v199 = v197 - int32(96)
	m.G0 = v199
	v202 = v48 + int32(80)
	v203 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v203
	v208 = v48 + int32(72)
	v209 = int64(4607182418800017408)
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = v209
	v212 = v48 + int32(88)
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = v209
	if v196 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v94 = int32(0)
	goto L18
L18:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v89+v94<<(uint(int32(2))%32))))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 != v63 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L15
L20:
	;
	v149 = v94 + int32(1)
	if v86 != v149 {
		v94 = v149
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v142 != v71 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	if v144 != v75 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+12)))
	if v146 == v78 {
		v671 = v139
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	goto L19
L26:
	;
	m.G0 = v199 + int32(96)
	v628 = int32(_a_F_initial_cost_mergejoin_0)
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[0]))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[0])) = v631
	v634 = F_palloc(m, int32(48))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L32
	} else {
		goto L145
	}
L27:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v217 != int32(17) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v196)+28))
	if v220 == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v223 < int32(2) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v227 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v196)+24))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	F_examine_variable(m, l0, v232, int32(0), v199-int32(-64))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return
L33:
	;
	F_examine_variable(m, l0, v227, int32(0), v199+int32(32))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v243 = F_get_opfamily_method(m, v63)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	F_get_op_opfamily_properties(m, v231, v63, int32(0), v199+int32(28), v199+int32(24), v199+int32(20))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	switch v75 - int32(1) {
	case 0:
		goto L40
	default:
		goto L37
	case 4:
		goto L39
	}
L37:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	if v602 != 0 {
		goto L139
	} else {
		goto L140
	}
L38:
	;
	v359 = int32(0)
	if base.B2i32(v352 == v359)|base.B2i32(v350 == v359)|(base.B2i32(v353 == v359)|base.B2i32(v349 == v359))|(base.B2i32(v351 == v359)|base.B2i32(v357 == v359)|(base.B2i32(v354 == v359)|base.B2i32(v355 == v359))) != 0 {
		goto L37
	} else {
		goto L71
	}
L39:
	;
	v295 = int32(1)
	v298 = F_IndexAmTranslateCompareType(m, v295, v243, v63, v295)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L32
	} else {
		goto L54
	}
L40:
	;
	v256 = int32(1)
	v258 = F_IndexAmTranslateCompareType(m, v256, v243, v63, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v262 = F_IndexAmTranslateCompareType(m, int32(2), v243, v63, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v264 = base.I32_extend16_s(v262)
	v265 = base.I32_extend16_s(v258)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	if v266 == v267 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v269 = F_get_opfamily_member(m, v63, v266, v266, v265)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L32
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v275 = F_get_opfamily_member(m, v63, v266, v267, v265)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L32
	} else {
		goto L48
	}
L46:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v273 = F_get_opfamily_member(m, v63, v271, v272, v264)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v349 = v269
	v350 = v269
	v351 = v269
	v352 = v269
	v353 = v269
	v354 = v269
	v355 = v273
	v356 = v10
	v357 = v273
	goto L38
L48:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v279 = F_get_opfamily_member(m, v63, v277, v278, v264)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L32
	} else {
		goto L49
	}
L49:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v282 = F_get_opfamily_member(m, v63, v281, v281, v265)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v285 = F_get_opfamily_member(m, v63, v284, v284, v265)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L32
	} else {
		goto L51
	}
L51:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v289 = F_get_opfamily_member(m, v63, v287, v288, v265)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L32
	} else {
		goto L52
	}
L52:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v293 = F_get_opfamily_member(m, v63, v291, v292, v264)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v349 = v285
	v350 = v285
	v351 = v275
	v352 = v282
	v353 = v282
	v354 = v289
	v355 = v293
	v356 = v10
	v357 = v279
	goto L38
L54:
	;
	v302 = F_IndexAmTranslateCompareType(m, int32(5), v243, v63, int32(1))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L32
	} else {
		goto L55
	}
L55:
	;
	v304 = base.I32_extend16_s(v302)
	v307 = F_IndexAmTranslateCompareType(m, int32(4), v243, v63, int32(1))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L32
	} else {
		goto L56
	}
L56:
	;
	v309 = base.I32_extend16_s(v307)
	v310 = base.I32_extend16_s(v298)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	if v311 == v312 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v314 = F_get_opfamily_member(m, v63, v311, v311, v304)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L32
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v323 = F_get_opfamily_member(m, v63, v311, v312, v304)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L32
	} else {
		goto L63
	}
L60:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v318 = F_get_opfamily_member(m, v63, v316, v317, v309)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v321 = F_get_opfamily_member(m, v63, v320, v320, v310)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L32
	} else {
		goto L62
	}
L62:
	;
	v349 = v321
	v350 = v314
	v351 = v314
	v352 = v314
	v353 = v321
	v354 = v314
	v355 = v318
	v356 = v295
	v357 = v318
	goto L38
L63:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v327 = F_get_opfamily_member(m, v63, v325, v326, v309)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L32
	} else {
		goto L64
	}
L64:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v330 = F_get_opfamily_member(m, v63, v329, v329, v304)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L32
	} else {
		goto L65
	}
L65:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v333 = F_get_opfamily_member(m, v63, v332, v332, v304)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L32
	} else {
		goto L66
	}
L66:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v336 = F_get_opfamily_member(m, v63, v335, v335, v310)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L32
	} else {
		goto L67
	}
L67:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v339 = F_get_opfamily_member(m, v63, v338, v338, v310)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L32
	} else {
		goto L68
	}
L68:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v343 = F_get_opfamily_member(m, v63, v341, v342, v304)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L32
	} else {
		goto L69
	}
L69:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v347 = F_get_opfamily_member(m, v63, v345, v346, v309)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L32
	} else {
		goto L70
	}
L70:
	;
	v349 = v339
	v350 = v333
	v351 = v323
	v352 = v330
	v353 = v336
	v354 = v343
	v355 = v347
	v356 = v295
	v357 = v327
	goto L38
L71:
	;
	if v356 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v427 = F_scalarineqsel(m, l0, v357, v356, int32(1), v230, v199-int32(-64), v425, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L32
	} else {
		goto L84
	}
L73:
	;
	v390 = F_get_variable_range(m, v199-int32(-64), v353, v230, v199+int32(16), v199+int32(12))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L32
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v408 = F_get_variable_range(m, v199-int32(-64), v353, v230, v199+int32(12), v199+int32(16))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L32
	} else {
		goto L80
	}
L76:
	;
	if v390 == int32(0) {
		goto L37
	} else {
		goto L77
	}
L77:
	;
	v400 = F_get_variable_range(m, v199+int32(32), v349, v230, v199+int32(8), v199+int32(4))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L32
	} else {
		goto L78
	}
L78:
	;
	if v400 != 0 {
		goto L72
	} else {
		goto L79
	}
L79:
	;
	goto L37
L80:
	;
	if v408 == int32(0) {
		goto L37
	} else {
		goto L81
	}
L81:
	;
	v418 = F_get_variable_range(m, v199+int32(32), v349, v230, v199+int32(4), v199+int32(8))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L32
	} else {
		goto L82
	}
L82:
	;
	if v418 == int32(0) {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	goto L72
L84:
	;
	if base.F64_ne(v427, float64(0.3333333333333333)) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v212))) = v427
	goto L87
L86:
	;
	goto L87
L87:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v437 = F_scalarineqsel(m, l0, v355, v356, int32(1), v230, v199+int32(32), v435, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L32
	} else {
		goto L89
	}
L88:
	;
	v444 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.F64_gt(v444, v443) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	if base.F64_eq(v437, float64(0.3333333333333333)) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v441 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	v443 = v441
	goto L88
L91:
	;
	goto L92
L92:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v208))) = v437
	v443 = v437
	goto L88
L93:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v458 = F_scalarineqsel(m, l0, v351, v356, int32(0), v230, v199-int32(-64), v456, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L32
	} else {
		goto L98
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = int64(4607182418800017408)
	if base.F64_gt(v443, v444) != 0 {
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = int64(4607182418800017408)
	goto L93
L97:
	;
	goto L96
L98:
	;
	if base.F64_ne(v458, float64(0.3333333333333333)) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v458
	goto L101
L100:
	;
	goto L101
L101:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v468 = F_scalarineqsel(m, l0, v354, v356, int32(0), v230, v199+int32(32), v466, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L32
	} else {
		goto L103
	}
L102:
	;
	v475 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	if base.F64_lt(v475, v474) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	if base.F64_eq(v468, float64(0.3333333333333333)) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v472 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	v474 = v472
	goto L102
L105:
	;
	goto L106
L106:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v202))) = v468
	v474 = v468
	goto L102
L107:
	;
	if v78 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = int64(0)
	if base.F64_lt(v474, v475) != 0 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(0)
	goto L107
L111:
	;
	goto L110
L112:
	;
	v574 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v575 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.F64_ge(v574, v575) != 0 {
		goto L135
	} else {
		goto L136
	}
L113:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	if v486 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	if v530 == int32(0) {
		goto L112
	} else {
		goto L125
	}
L115:
	;
	v489 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+22)))
	v492 = v490 + v491
	v493 = *(*float32)(unsafe.Add(mBase, uint32(v492)+8))
	v495 = base.F64_add(v489, base.F64_promote_f32(v493))
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v495
	v497 = float64(0)
	if base.F64_lt(v495, v497) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v510 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	v511 = *(*float32)(unsafe.Add(mBase, uint32(v492)+8))
	v513 = base.F64_add(v510, base.F64_promote_f32(v511))
	*(*float64)(unsafe.Add(mBase, uint32(v212))) = v513
	v515 = float64(0)
	if base.F64_lt(v513, v515) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v502 = float64(1)
	if base.F64_gt(v495, v502) == int32(0) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	v507 = v497
	goto L119
L119:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v507
	goto L116
L120:
	;
	v507 = v502
	goto L119
L121:
	;
	v520 = float64(1)
	if base.F64_gt(v513, v520) == int32(0) {
		goto L114
	} else {
		goto L124
	}
L122:
	;
	v525 = v515
	goto L123
L123:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v212))) = v525
	goto L114
L124:
	;
	v525 = v520
	goto L123
L125:
	;
	v533 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530)+16))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+22)))
	v536 = v534 + v535
	v537 = *(*float32)(unsafe.Add(mBase, uint32(v536)+8))
	v539 = base.F64_add(v533, base.F64_promote_f32(v537))
	*(*float64)(unsafe.Add(mBase, uint32(v202))) = v539
	v541 = float64(0)
	if base.F64_lt(v539, v541) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v554 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	v555 = *(*float32)(unsafe.Add(mBase, uint32(v536)+8))
	v557 = base.F64_add(v554, base.F64_promote_f32(v555))
	*(*float64)(unsafe.Add(mBase, uint32(v208))) = v557
	v559 = float64(0)
	if base.F64_lt(v557, v559) == int32(0) {
		goto L131
	} else {
		goto L132
	}
L127:
	;
	v546 = float64(1)
	if base.F64_gt(v539, v546) == int32(0) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	v551 = v541
	goto L129
L129:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v202))) = v551
	goto L126
L130:
	;
	v551 = v546
	goto L129
L131:
	;
	v564 = float64(1)
	if base.F64_gt(v557, v564) == int32(0) {
		goto L112
	} else {
		goto L134
	}
L132:
	;
	v569 = v559
	goto L133
L133:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v208))) = v569
	goto L112
L134:
	;
	v569 = v564
	goto L133
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = int64(4607182418800017408)
	goto L137
L136:
	;
	goto L137
L137:
	;
	v581 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	v582 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	if base.F64_ge(v581, v582) == int32(0) {
		goto L37
	} else {
		goto L138
	}
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = int64(4607182418800017408)
	goto L37
L139:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	m.T0[v603].(func(*base.Module, int32))(m, v602)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L32
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	if v606 == int32(0) {
		goto L26
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v199)+44))
	m.T0[v609].(func(*base.Module, int32))(m, v606)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L32
	} else {
		goto L144
	}
L144:
	;
	goto L26
L145:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v634)+4)) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v634)+8)) = v641
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v634)+12)) = uint8(v643)
	v645 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	*(*float64)(unsafe.Add(mBase, uint32(v634)+16)) = v645
	v647 = *(*float64)(unsafe.Add(mBase, uint32(v48)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v634)+24)) = v647
	v649 = *(*float64)(unsafe.Add(mBase, uint32(v48)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v634)+32)) = v649
	v651 = *(*float64)(unsafe.Add(mBase, uint32(v48)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v634)+40)) = v651
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v82)+116))
	v654 = F_lappend(m, v653, v634)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L32
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+116)) = v654
	*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[0])) = v629
	v671 = v634
	goto L14
L147:
	;
	if v762 != 0 {
		goto L161
	} else {
		goto L162
	}
L148:
	;
	v762 = int32(1)
	goto L147
L149:
	;
	goto L150
L150:
	;
	if v708 == int32(0) {
		v755 = v709
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v762 = v755
	goto L147
L152:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v719 < v718 {
		v755 = v709
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v721 = int32(1)
	if v718 <= v721 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v724 = v721
	goto L156
L155:
	;
	v724 = v718
	goto L156
L156:
	;
	v725 = int32(8)
	v730 = int32(0)
	goto L157
L157:
	;
	v737 = v730 << (uint(int32(2)) % 32)
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v706+v725+v737)))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v708+v725+v737)))
	v744 = v739 & (v741 ^ int32(-1))
	v746 = base.B2i32(v744 == int32(0))
	if v744 != 0 {
		v755 = v746
		goto L151
	} else {
		goto L159
	}
L158:
	;
	v755 = v746
	goto L151
L159:
	;
	v748 = v730 + int32(1)
	if v748 != v724 {
		v730 = v748
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v763 = int32(32)
	goto L163
L162:
	;
	v763 = int32(16)
	goto L163
L163:
	;
	v765 = *(*float64)(unsafe.Add(mBase, uint32(v671+v763)))
	if v762 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v768 = int32(40)
	goto L166
L165:
	;
	v768 = int32(24)
	goto L166
L166:
	;
	v770 = *(*float64)(unsafe.Add(mBase, uint32(v671+v768)))
	v772 = l2 & int32(-5)
	if v772 == int32(1) {
		v819 = v770
		v820 = v765
		v822 = v52
		v825 = v29
		goto L2
	} else {
		goto L167
	}
L167:
	;
	if v762 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v777 = int32(24)
	goto L170
L169:
	;
	v777 = int32(40)
	goto L170
L170:
	;
	v779 = *(*float64)(unsafe.Add(mBase, uint32(v671+v777)))
	if v762 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v782 = int32(16)
	goto L173
L172:
	;
	v782 = int32(32)
	goto L173
L173:
	;
	v784 = *(*float64)(unsafe.Add(mBase, uint32(v671+v782)))
	v787 = base.B2i32(v772 == int32(3))
	if v772 == int32(3) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v788 = float64(1)
	goto L176
L175:
	;
	v788 = v770
	goto L176
L176:
	;
	if v772 == int32(3) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v790 = float64(0)
	goto L179
L178:
	;
	v790 = v765
	goto L179
L179:
	;
	v819 = v788
	v820 = v790
	v822 = v779
	v825 = v784
	goto L2
L180:
	;
	v839 = float64(1)
	goto L182
L181:
	;
	v839 = v50
	goto L182
L182:
	;
	if base.F64_le(v51, float64(0)) != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v844 = float64(1)
	goto L185
L184:
	;
	v844 = v51
	goto L185
L185:
	;
	v846 = float64(1e+100)
	v847 = base.F64_mul(v844, v822)
	if base.F64_gt(v847, v846) != 0 {
		v860 = v846
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v862 = base.F64_nearest(base.F64_mul(v844, v825))
	v863 = base.F64_mul(v839, v819)
	if base.F64_gt(v863, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v863)&int64(9223372036854775807))) != 0 {
		v876 = float64(1e+100)
		goto L190
	} else {
		goto L191
	}
L187:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v847)&int64(9223372036854775807)) {
		v860 = float64(1e+100)
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v856 = float64(1)
	if base.F64_le(v847, v856) != 0 {
		v860 = v856
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v860 = base.F64_nearest(v847)
	goto L186
L190:
	;
	v877 = base.F64_nearest(base.F64_mul(v839, v820))
	v878 = base.F64_div(v862, v844)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[1])))
	if l6 != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v872 = float64(1)
	if base.F64_le(v863, v872) != 0 {
		v876 = v872
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v876 = base.F64_nearest(v863)
	goto L190
L193:
	;
	v1024 = base.F64_div(v877, v839)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	if l7 != 0 {
		goto L223
	} else {
		goto L224
	}
L194:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	if l8 <= int32(0) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	goto L196
L196:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	v1007 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v1008 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v1009 = base.F64_sub(v1007, v1008)
	v1015 = v1006
	v1020 = v1009
	v1023 = base.F64_add(base.F64_mul(v1009, v878), base.F64_add(v1008, float64(0)))
	goto L193
L197:
	;
	v1001 = base.F64_sub(v1000, v995)
	v1015 = v993
	v1020 = v1001
	v1023 = base.F64_add(base.F64_mul(v1001, v878), base.F64_add(v995, float64(0)))
	goto L193
L198:
	;
	v902 = float64(2)
	if base.F64_lt(v844, v902) != 0 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[2])))
	if v885&int32(1) == int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v890 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v891 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v892)+32))
	v895 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3]))
	F_cost_incremental_sort(m, v48, l0, l6, l8, v881, v890, v891, v844, v893, v895, float64(-1))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L32
	} else {
		goto L201
	}
L201:
	;
	v899 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	v901 = *(*float64)(unsafe.Add(mBase, uint32(v48)+56))
	v993 = v900
	v995 = v899
	v1000 = v901
	goto L197
L202:
	;
	v905 = v902
	goto L204
L203:
	;
	v905 = v844
	goto L204
L204:
	;
	v907 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[4]))
	v911 = base.F64_mul(v905, base.F64_add(base.F64_add(v907, v907), float64(0)))
	v912 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+32))
	v922 = base.F64_mul(v844, base.F64_convert_i32_u((v914+int32(7))&int32(-8)+int32(24)))
	v924 = int64(*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3])))
	v926 = v924 << (uint(int64(10)) % 64)
	v927 = base.F64_convert_i64_s(v926)
	if base.F64_gt(v922, v927) != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[5])))
	v990 = base.F64_add(v912, v983)
	v993 = v881 + (v986 ^ int32(1))
	v995 = v990
	v1000 = base.F64_add(v990, base.F64_mul(v905, v982))
	goto L197
L206:
	;
	v929 = F_log(m, v905)
	mBase = m.M
	v935 = base.F64_ceil(base.F64_mul(v922, float64(0.0001220703125)))
	v937 = base.F64_div(v922, v927)
	v940 = int32(6)
	v942 = base.I64_div_s(v926, int64(278528))
	v943 = base.I32_wrap_i64(v942)
	if v943 <= v940 {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	goto L208
L208:
	;
	v972 = base.F64_add(v905, v905)
	if base.F64_lt(v972, v905) != 0 {
		goto L219
	} else {
		goto L220
	}
L209:
	;
	v950 = base.F64_convert_i32_s(v949)
	if base.F64_gt(v937, v950) != 0 {
		goto L216
	} else {
		goto L217
	}
L210:
	;
	v946 = v940
	goto L212
L211:
	;
	v946 = v943
	goto L212
L212:
	;
	if int32(500) <= v946 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v949 = int32(500)
	goto L215
L214:
	;
	v949 = v946
	goto L215
L215:
	;
	goto L209
L216:
	;
	v952 = F_log(m, v937)
	mBase = m.M
	v953 = F_log(m, v950)
	mBase = m.M
	v957 = base.F64_ceil(base.F64_div(v952, v953))
	goto L218
L217:
	;
	v957 = float64(1)
	goto L218
L218:
	;
	v960 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[6]))
	v964 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[7]))
	v971 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[4]))
	v982 = v971
	v983 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v935, v935), v957), base.F64_add(base.F64_mul(v960, float64(0.75)), base.F64_mul(v964, float64(0.25)))), base.F64_mul(base.F64_div(v929, float64(0.693147180559945)), v911))
	goto L205
L219:
	;
	v974 = F_log(m, v972)
	mBase = m.M
	v982 = v907
	v983 = base.F64_mul(base.F64_div(v974, float64(0.693147180559945)), v911)
	goto L205
L220:
	;
	goto L221
L221:
	;
	v978 = F_log(m, v905)
	mBase = m.M
	v982 = v907
	v983 = base.F64_mul(base.F64_div(v978, float64(0.693147180559945)), v911)
	goto L205
L222:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+72)) = v877
	*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = v862
	*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v876
	*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = v860
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1134
	v1142 = base.F64_mul(base.F64_sub(base.F64_div(v876, v839), v1024), v1128)
	*(*float64)(unsafe.Add(mBase, uint32(l1)+32)) = v1142
	v1148 = base.F64_add(base.F64_mul(v1020, base.F64_sub(base.F64_div(v860, v844), v878)), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v1148
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1015 + (v880 ^ int32(1)) + v1127
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v1142, base.F64_add(v1148, v1134))
	m.G0 = v48 + int32(96)
	return
L223:
	;
	v1026 = float64(2)
	if base.F64_lt(v839, v1026) != 0 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v1121 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v1122 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v1123 = base.F64_sub(v1121, v1122)
	v1127 = v1025
	v1128 = v1123
	v1134 = base.F64_add(base.F64_mul(v1123, v1024), base.F64_add(v1023, v1122))
	goto L222
L226:
	;
	v1029 = v1026
	goto L228
L227:
	;
	v1029 = v839
	goto L228
L228:
	;
	v1031 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[4]))
	v1035 = base.F64_mul(v1029, base.F64_add(base.F64_add(v1031, v1031), float64(0)))
	v1036 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+32))
	v1046 = base.F64_mul(v839, base.F64_convert_i32_u((v1038+int32(7))&int32(-8)+int32(24)))
	v1048 = int64(*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3])))
	v1050 = v1048 << (uint(int64(10)) % 64)
	v1051 = base.F64_convert_i64_s(v1050)
	if base.F64_gt(v1046, v1051) != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[5])))
	v1114 = base.F64_add(v1036, v1107)
	v1117 = base.F64_sub(base.F64_add(v1114, base.F64_mul(v1029, v1106)), v1114)
	v1127 = v1025 + (v1110 ^ int32(1))
	v1128 = v1117
	v1134 = base.F64_add(base.F64_mul(v1117, v1024), base.F64_add(v1023, v1114))
	goto L222
L230:
	;
	v1053 = F_log(m, v1029)
	mBase = m.M
	v1059 = base.F64_ceil(base.F64_mul(v1046, float64(0.0001220703125)))
	v1061 = base.F64_div(v1046, v1051)
	v1064 = int32(6)
	v1066 = base.I64_div_s(v1050, int64(278528))
	v1067 = base.I32_wrap_i64(v1066)
	if v1067 <= v1064 {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	goto L232
L232:
	;
	v1096 = base.F64_add(v1029, v1029)
	if base.F64_lt(v1096, v1029) != 0 {
		goto L243
	} else {
		goto L244
	}
L233:
	;
	v1074 = base.F64_convert_i32_s(v1073)
	if base.F64_gt(v1061, v1074) != 0 {
		goto L240
	} else {
		goto L241
	}
L234:
	;
	v1070 = v1064
	goto L236
L235:
	;
	v1070 = v1067
	goto L236
L236:
	;
	if int32(500) <= v1070 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1073 = int32(500)
	goto L239
L238:
	;
	v1073 = v1070
	goto L239
L239:
	;
	goto L233
L240:
	;
	v1076 = F_log(m, v1061)
	mBase = m.M
	v1077 = F_log(m, v1074)
	mBase = m.M
	v1081 = base.F64_ceil(base.F64_div(v1076, v1077))
	goto L242
L241:
	;
	v1081 = float64(1)
	goto L242
L242:
	;
	v1084 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[6]))
	v1088 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[7]))
	v1095 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[4]))
	v1106 = v1095
	v1107 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v1059, v1059), v1081), base.F64_add(base.F64_mul(v1084, float64(0.75)), base.F64_mul(v1088, float64(0.25)))), base.F64_mul(base.F64_div(v1053, float64(0.693147180559945)), v1035))
	goto L229
L243:
	;
	v1098 = F_log(m, v1096)
	mBase = m.M
	v1106 = v1031
	v1107 = base.F64_mul(base.F64_div(v1098, float64(0.693147180559945)), v1035)
	goto L229
L244:
	;
	goto L245
L245:
	;
	v1102 = F_log(m, v1029)
	mBase = m.M
	v1106 = v1031
	v1107 = base.F64_mul(base.F64_div(v1102, float64(0.693147180559945)), v1035)
	goto L229
L246:
	;
	F_errmsg_internal(m, int32(_a_F_initial_cost_mergejoin_1), int32(0))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L32
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_initial_cost_mergejoin_2), int32(3615), int32(_a_F_initial_cost_mergejoin_3))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L32
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_initial_cost_nestloop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 float64
	_ = v58
	var v60 int32
	_ = v60
	var v68 float64
	_ = v68
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v97 float64
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v122 float64
	_ = v122
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v151 float64
	_ = v151
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 float64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 float64
	_ = v188
	var v198 int32
	_ = v198
	var v208 float64
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 float64
	_ = v213
	var v216 float64
	_ = v216
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v222 float64
	_ = v222
	var v226 float64
	_ = v226
	var v229 float64
	_ = v229
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v244 float64
	_ = v244
	var v247 float64
	_ = v247
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v264 float64
	_ = v264
	var v265 float64
	_ = v265
	var v281 float64
	_ = v281
	var v283 float64
	_ = v283
	var v284 float64
	_ = v284
	var v287 float64
	_ = v287
	var v291 float64
	_ = v291
	var v292 float64
	_ = v292
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v295 float64
	_ = v295
	var v300 int32
	_ = v300
	var v305 float64
	_ = v305
	var v312 float64
	_ = v312
	var v316 float64
	_ = v316
	v7 = float64(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[0])))
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v33 - int32(348) {
	case 0:
		goto L7
	default:
		goto L2
	case 3, 5:
		goto L5
	case 11:
		goto L6
	case 12, 14:
		goto L4
	case 13:
		goto L3
	}
L1:
	;
	v281 = base.F64_add(v32, float64(-1))
	v283 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v284 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v287 = base.F64_add(base.F64_sub(v283, v284), float64(0))
	if base.F64_gt(v32, float64(1)) != 0 {
		goto L41
	} else {
		goto L42
	}
L2:
	;
	v255 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v264 = v255
	v265 = v256
	goto L1
L3:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+72))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v103)+32))
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l4)+88))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v103)+56))
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v103)+48))
	v112 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[1]))
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[2]))
	v118 = base.F64_mul(base.F64_mul(v112, base.F64_convert_i32_s(v114)), float64(1024))
	v119 = float64(4.294967295e+09)
	if base.F64_lt(v118, v119) != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v75 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[3]))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v77 = base.F64_mul(v75, v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+32))
	v87 = base.F64_mul(v76, base.F64_convert_i32_u((v79+int32(7))&int32(-8)+int32(24)))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[2]))
	if base.F64_gt(v87, base.F64_convert_i32_u(v89<<(uint(int32(10))%32))) == int32(0) {
		v264 = v77
		v265 = v7
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v46 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[4]))
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v48 = base.F64_mul(v46, v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
	v58 = base.F64_mul(v47, base.F64_convert_i32_u((v50+int32(7))&int32(-8)+int32(24)))
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[2]))
	if base.F64_gt(v58, base.F64_convert_i32_u(v60<<(uint(int32(10))%32))) == int32(0) {
		v264 = v48
		v265 = v7
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+100))
	if v39 != int32(1) {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v264 = base.F64_sub(v36, v37)
	v265 = v7
	goto L1
L8:
	;
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v264 = base.F64_sub(v42, v43)
	v265 = v7
	goto L1
L9:
	;
	v68 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[5]))
	v264 = base.F64_add(base.F64_mul(v68, base.F64_ceil(base.F64_mul(v58, float64(0.0001220703125)))), v48)
	v265 = v7
	goto L1
L10:
	;
	v97 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[5]))
	v264 = base.F64_add(base.F64_mul(v97, base.F64_ceil(base.F64_mul(v87, float64(0.0001220703125)))), v77)
	v265 = v7
	goto L1
L11:
	;
	v136 = base.F64_add(base.F64_add(base.F64_mul(v106, float64(8)), float64(28)), base.F64_mul(v106, base.F64_convert_i32_u((v105+int32(7))&int32(-8)+int32(24))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	if v137 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v122 = v118
	goto L14
L13:
	;
	v122 = v119
	goto L14
L14:
	;
	goto L11
L15:
	;
	v208 = F_estimate_num_groups(m, l0, v198, v107, int32(0), v26+int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L22
	} else {
		goto L25
	}
L16:
	;
	v188 = v136
	v198 = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v141 <= int32(0) {
		v188 = v136
		v198 = v137
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v151 = v136
	v162 = int32(0)
	goto L20
L20:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v162<<(uint(int32(2))%32))))
	v173 = F_get_expr_width(m, l0, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v188 = v176
	v198 = v181
	goto L15
L22:
	;
	return
L23:
	;
	v176 = base.F64_add(v151, base.F64_convert_i32_s(v173))
	v178 = v162 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v178 < v179 {
		v151 = v176
		v162 = v178
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v210&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v213 = v107
	goto L28
L27:
	;
	v213 = v208
	goto L28
L28:
	;
	v216 = base.F64_floor(base.F64_div(base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v122)), v188))
	if base.F64_gt(v216, v213) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v218 = v213
	goto L31
L30:
	;
	v218 = v216
	goto L31
L31:
	;
	v219 = float64(4.294967295e+09)
	if base.F64_lt(v218, v219) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v222 = v218
	goto L34
L33:
	;
	v222 = v219
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+96)) = base.I32_trunc_sat_f64_u(v222)
	v226 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[3]))
	v229 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[4]))
	v234 = base.F64_lt(v216, v213)
	if v234 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v235 = v216
	goto L37
L36:
	;
	v235 = v213
	goto L37
L37:
	;
	v237 = base.F64_sub(float64(1), base.F64_div(v235, v213))
	if v234 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v244 = v213
	goto L40
L39:
	;
	v244 = v216
	goto L40
L40:
	;
	v247 = base.F64_sub(float64(1), base.F64_mul(base.F64_div(base.F64_sub(v107, v213), v107), base.F64_div(v216, v244)))
	v264 = base.F64_add(base.F64_add(base.F64_mul(v226, v106), v229), base.F64_add(base.F64_mul(base.F64_mul(base.F64_div(v226, float64(10)), v237), v106), base.F64_add(base.F64_mul(v229, v237), base.F64_add(base.F64_mul(v108, v247), v226))))
	v265 = base.F64_add(v229, base.F64_mul(v109, v247))
	goto L1
L41:
	;
	v291 = base.F64_add(base.F64_mul(v281, v265), v287)
	goto L43
L42:
	;
	v291 = v287
	goto L43
L43:
	;
	v292 = base.F64_sub(v264, v265)
	v293 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v294 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v295 = base.F64_sub(v293, v294)
	if l2&int32(-2) != int32(4) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v312
	v316 = base.F64_add(base.F64_add(v284, v294), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v316
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v316, v312)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29 + (v31 ^ int32(1)) + v28
	m.G0 = v26 + int32(16)
	return
L45:
	;
	v305 = base.F64_add(v295, v291)
	if base.F64_gt(v32, float64(1)) == int32(0) {
		v312 = v305
		goto L44
	} else {
		goto L50
	}
L46:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v300 != int32(1) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = v292
	*(*float64)(unsafe.Add(mBase, uint32(l1)+32)) = v295
	v312 = v291
	goto L44
L49:
	;
	goto L48
L50:
	;
	v312 = base.F64_add(base.F64_mul(v281, v292), v305)
	goto L44
}
func F_initialize_acl(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_acl[0]))
	if v2 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_acl[1]))
		v8 = F_GetSysCacheHashValue(m, int32(21), v6, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_initialize_acl[2])) = v8
			F_CacheRegisterSyscacheCallback(m, int32(9), int32(1239), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(11), int32(1239), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(21), int32(1239), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_inject_projection_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 float64
	_ = v22
	var v24 float64
	_ = v24
	var v26 float64
	_ = v26
	var v28 int32
	_ = v28
	v3 = l2
	v6 = F_palloc0(m, int32(80))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(331)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v20
		v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v22
		v24 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+16)) = v24
		v26 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+24)) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+37)) = uint8(v3)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+36)) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v28
		return v6
	}
}
func F_inline_cte_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v3 = int32(0)
	if l0 == v3 {
		v92 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v92
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != int32(101) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v89 = F_expression_tree_walker_impl(m, l0, int32(844), l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L25
	}
L4:
	;
	if v7 != int32(67) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 != int32(6) {
		v92 = v3
		goto L1
	} else {
		goto L10
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v12 + int32(1)
	v18 = F_query_tree_walker_impl(m, l0, int32(844), l1, int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 - int32(1)
	return int32(0)
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if base.B2i32(v35 == int32(0))|base.B2i32(v35 != v38) != 0 {
		v56 = v35
		v57 = v38
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v56-v57 != 0 {
		v92 = v3
		goto L1
	} else {
		goto L18
	}
L12:
	;
	goto L11
L13:
	;
	v41 = v31
	v42 = v32
	goto L14
L14:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v46 == int32(0) {
		v56 = v46
		v57 = v45
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v56 = v46
	v57 = v45
	goto L12
L16:
	;
	v49 = int32(1)
	if v46 == v45 {
		v41 = v41 + v49
		v42 = v42 + v49
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v59 != v60 {
		v92 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v63 = l0 + int32(84)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v65 = F_copyObjectImpl(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v67 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_IncrementVarSublevelsUp(m, v65, v67, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v73
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)) = uint8(v73)
	*(*int64)(unsafe.Add(mBase, uint32(v63))) = v75
	return v73
L24:
	;
	goto L23
L25:
	;
	v92 = v89
	goto L1
}
func F_insertStatEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var __phi242 int32
	_ = __phi242
	var v244 int32
	_ = v244
	var __phi244 int32
	_ = __phi244
	var v249 int32
	_ = v249
	var __phi249 int32
	_ = __phi249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	v5 = int32(0)
	v17 = int32(1)
	v19 = l2 + int32(8)
	v22 = v19 + l3<<(uint(int32(2))%32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = v23 & v17
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v27 == v5 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	if v26 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	if v184 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	if v25 == int32(0) {
		v202 = v17
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v36 = int32(1)
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v32<<(uint(int32(2))%32)+(int32(base.Ui32(v23)>>(uint(v36)%32))&int32(2047)+int32(base.Ui32(v23)>>(uint(int32(12))%32))+v36)&int32(_a_F_insertStatEntry_0)))))
	v184 = v48
	goto L3
L8:
	;
	v51 = int32(1)
	v61 = (int32(base.Ui32(v23)>>(uint(v51)%32))&int32(2047) + int32(base.Ui32(v23)>>(uint(int32(12))%32)) + v51) & int32(_a_F_insertStatEntry_0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v64 = v62 << (uint(int32(2)) % 32)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+(v19+v64)))))
	if v67 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v71 = v67 & int32(3)
	v75 = l2 + v64 + v61 + int32(10)
	if base.Ui32(v67) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v152 = v136
	v156 = v140
	v157 = v5
	goto L18
L11:
	;
	v136 = v75
	v140 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v85 = v75
	v89 = int32(0)
	v92 = v5
	goto L14
L14:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85))))
	v99 = int32(14)
	v102 = int32(1)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+2)))
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+6)))
	v125 = int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v98)>>(uint(v99)%32)))%32))&v102 + v89 + int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v105)>>(uint(v99)%32)))%32))&v102 + int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v112)>>(uint(v99)%32)))%32))&v102 + int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v119)>>(uint(v99)%32)))%32))&v102
	v127 = v85 + int32(8)
	v129 = v92 + int32(4)
	if v129 != v67&int32(_a_F_insertStatEntry_1) {
		v85 = v127
		v89 = v125
		v92 = v129
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if v71 == int32(0) {
		v184 = v125
		goto L3
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v136 = v127
	v140 = v125
	goto L10
L18:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v169 = int32(1)
	v171 = int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v165)>>(uint(int32(14))%32)))%32))&v169 + v156
	v175 = v157 + v169
	if v175 != v71 {
		v152 = v152 + int32(2)
		v156 = v171
		v157 = v175
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v184 = v171
	goto L3
L20:
	;
	goto L19
L21:
	;
	v202 = v184
	goto L2
L22:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v367) < base.Ui32(v358) {
		goto L71
	} else {
		goto L72
	}
L23:
	;
	v213 = int32(0)
	v214 = int32(1)
	v353 = v213
	v354 = v213
	v356 = v214
	v358 = v214
	v366 = v213
	goto L22
L24:
	;
	goto L25
L25:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v224 = v19 + v218<<(uint(int32(2))%32) + int32(base.Ui32(v23)>>(uint(int32(12))%32))
	v230 = int32(base.Ui32(v23)>>(uint(int32(1))%32)) & int32(2047)
	if v230 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v231 = int32(-1)
	goto L28
L27:
	;
	v231 = int32(0)
	goto L28
L28:
	;
	__phi242 = int32(1)
	__phi244 = int32(0)
	__phi249 = v26
	v242 = __phi242
	v244 = __phi244
	v249 = __phi249
	goto L29
L29:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v250 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v353 = v341
	v354 = v342
	v356 = v344
	v358 = v345
	v366 = int32(base.Ui32(v343) >> (uint(int32(31)) % 32))
	goto L22
L31:
	;
	goto L30
L32:
	;
	v330 = int32(1)
	v332 = v242 + v330
	v333 = int32(0)
	if v329 < v333 {
		goto L67
	} else {
		goto L68
	}
L33:
	;
	v257 = v249 + int32(20)
	if base.Ui32(v250) < base.Ui32(v230) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v230 != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v253 = v231
	goto L36
L36:
	;
	if v253 != 0 {
		v329 = v253
		goto L32
	} else {
		goto L38
	}
L37:
	;
	v253 = base.B2i32(int32(0) < v250)
	goto L36
L38:
	;
	v254 = int32(0)
	v341 = v244
	v342 = v249
	v343 = v254
	v344 = v254
	v345 = v242
	goto L31
L39:
	;
	v259 = v250
	goto L41
L40:
	;
	v259 = v230
	goto L41
L41:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v259) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	if v321 != 0 {
		v329 = v321
		goto L32
	} else {
		goto L60
	}
L43:
	;
	v321 = int32(0)
	goto L42
L44:
	;
	v295 = v290
	v296 = v291
	v297 = v292
	goto L54
L45:
	;
	if (v257|v224)&int32(3) != 0 {
		v290 = v257
		v291 = v224
		v292 = v259
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v283 = v257
	v284 = v224
	v285 = v259
	goto L47
L47:
	;
	if v285 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	v267 = v257
	v268 = v224
	v269 = v259
	goto L49
L49:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v272 != v273 {
		v290 = v267
		v291 = v268
		v292 = v269
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v283 = v278
	v284 = v276
	v285 = v280
	goto L47
L51:
	;
	v275 = int32(4)
	v276 = v268 + v275
	v278 = v267 + v275
	v280 = v269 - v275
	if base.Ui32(int32(3)) < base.Ui32(v280) {
		v267 = v278
		v268 = v276
		v269 = v280
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v290 = v283
	v291 = v284
	v292 = v285
	goto L44
L54:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v300 == v301 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v321 = v300 - v301
	goto L42
L56:
	;
	v303 = int32(1)
	v308 = v297 - v303
	if v308 != 0 {
		v295 = v295 + v303
		v296 = v296 + v303
		v297 = v308
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L43
L60:
	;
	if v250 == v230 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v323 = int32(0)
	v341 = v244
	v342 = v249
	v343 = v323
	v344 = v323
	v345 = v242
	goto L31
L62:
	;
	goto L63
L63:
	;
	if v250 < v230 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v328 = int32(-1)
	goto L66
L65:
	;
	v328 = int32(1)
	goto L66
L66:
	;
	v329 = v328
	goto L32
L67:
	;
	v338 = int32(8)
	goto L69
L68:
	;
	v338 = int32(12)
	goto L69
L69:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v249+v338)))
	if v340 != 0 {
		__phi242 = v332
		__phi244 = v249
		__phi249 = v340
		v242 = __phi242
		v244 = __phi244
		v249 = __phi249
		goto L29
	} else {
		goto L70
	}
L70:
	;
	v341 = v249
	v342 = v333
	v343 = v329
	v344 = v330
	v345 = v332
	goto L31
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v358
	goto L73
L72:
	;
	goto L73
L73:
	;
	if v356 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v377 = F_MemoryContextAlloc(m, l0, int32(base.Ui32(v370)>>(uint(int32(1))%32))&int32(2047)+int32(20))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v406 + int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+4)) = v410 + v202
	goto L1
L77:
	;
	return
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v377)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v202
	v382 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v388 = int32(base.Ui32(v384)>>(uint(v382)%32)) & int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+16)) = v388
	if v388 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	base.MemoryCopy(m, v377+int32(20), v19+v392<<(uint(int32(2))%32)+int32(base.Ui32(v396)>>(uint(int32(12))%32)), v388)
	goto L81
L80:
	;
	goto L81
L81:
	;
	if v353 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v377
	return
L83:
	;
	goto L84
L84:
	;
	if v366 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+8)) = v377
	return
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+12)) = v377
	return
}
func F_insert_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_replace_s(m, l0, l1, l2, l3, l4, v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			v31 = int32(-1)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if l1 <= v19 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21 + v19
			} else {
			}
			v24 = int32(0)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v25 < l1 {
				v31 = v24
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27 + v25
				v31 = v24
			}
		}
		m.G0 = v10 + int32(16)
		return v31
	}
}
func F_instantiate_empty_record_variable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 == int32(2249) {
		F_errstart_cold(m, int32(21), int32(_a_F_instantiate_empty_record_variable_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v18
				F_errmsg(m, int32(_a_F_instantiate_empty_record_variable_1), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errdetail(m, int32(_a_F_instantiate_empty_record_variable_2), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_instantiate_empty_record_variable_3), int32(_a_F_instantiate_empty_record_variable_4), int32(_a_F_instantiate_empty_record_variable_5))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_revalidate_rectypeid(m, l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v37 = F_make_expanded_record_from_typeid(m, v34, int32(-1), v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v37
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_int24mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v7 = v6 - v3
	if base.B2i32(int32(0) < v3) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int24mi_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int24mi_1), int32(1040), int32(_a_F_int24mi_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v7
	}
}
func F_int24mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 * v4
	v9 = base.I32_wrap_i64(v5)
	if base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) != v9>>(uint(int32(31))%32) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int24mul_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int24mul_1), int32(1054), int32(_a_F_int24mul_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v9
	}
}
func F_int24pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int24pl_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int24pl_1), int32(1026), int32(_a_F_int24pl_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v7
	}
}
func F_int28ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 <= v4)
}
func F_int28mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = int64(63)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v8) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v8 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8*(v12>>(uint(v9)%64)) + v8>>(uint(v9)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v29&v25 | v40<<(uint(v19)%64)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v51 != v52>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int28mul_0), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int28mul_1), int32(1150), int32(_a_F_int28mul_2))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v74 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v74
		}
	}
}
func F_int28pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v5 + v8
	if base.B2i32(v5 < int64(0)) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int28pl_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int28pl_1), int32(1122), int32(_a_F_int28pl_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int2mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 - v4
	v6 = base.I32_extend16_s(v5)
	if v6 != v5 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int2mi_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2mi_1), int32(958), int32(_a_F_int2mi_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v6
	}
}
func F_int2mod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = int32(_a_F_int2mod_0)
	v6 = v4 & v5
	if v6 != v5 {
		if v6 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2mod_1), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2mod_2), int32(1196), int32(_a_F_int2mod_3))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
			v13 = base.I32_rem_s(v11, base.I32_extend16_s(v4))
			v15 = v13
			return v15
		}
	} else {
		v15 = int32(0)
		return v15
	}
}
func F_int2mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 * v4
	v6 = base.I32_extend16_s(v5)
	if v6 != v5 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int2mul_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2mul_1), int32(972), int32(_a_F_int2mul_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v6
	}
}
func F_int2send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(2))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v19 = int32(8)
			v23 = v8<<(uint(v19)%32) | int32(base.Ui32(v8)>>(uint(v19)%32))
			*(*uint16)(unsafe.Add(mBase, uint32(v16+v17))) = uint16(v23)
			v25 = int32(2)
			v26 = v16 + v25
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v26
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v26 << (uint(v25) % 32)
			m.G0 = v6 + int32(16)
			return v29
		}
	}
}
func F_int2um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2&int32(_a_F_int2um_0) == int32(_a_F_int2um_1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int2um_2), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2um_3), int32(922), int32(_a_F_int2um_4))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v26 = int32(16)
		return (int32(0) - v2<<(uint(v26)%32)) >> (uint(v26) % 32)
	}
}
func F_int42mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 * v4
	v9 = base.I32_wrap_i64(v5)
	if base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) != v9>>(uint(int32(31))%32) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int42mul_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int42mul_1), int32(1115), int32(_a_F_int42mul_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v9
	}
}
func F_int48mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = int64(63)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v8) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v8 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8*(v12>>(uint(v9)%64)) + v8>>(uint(v9)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v29&v25 | v40<<(uint(v19)%64)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v51 != v52>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int48mul_0), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int48mul_1), int32(1008), int32(_a_F_int48mul_2))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v74 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v74
		}
	}
}
func F_int4gcd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var __phi28 int32
	_ = __phi28
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(31)
	v8 = v5 >> (uint(v7) % 32)
	v12 = v6 >> (uint(v7) % 32)
	v15 = base.B2i32(v12-(v12^v6) < v8-(v8^v5))
	if v12-(v12^v6) < v8-(v8^v5) {
		v16 = v5
	} else {
		v16 = v6
	}
	if v12-(v12^v6) < v8-(v8^v5) {
		v17 = v6
	} else {
		v17 = v5
	}
	if v17 == int32(-2147483648) {
		if v16&int32(2147483647) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int4gcd_0), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int4gcd_1), int32(1292), int32(_a_F_int4gcd_2))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if v16 != int32(-1) {
				__phi28 = v16
				__phi29 = v17
				v28 = __phi28
				v29 = __phi29
				for {
					v32 = base.I32_rem_s(v29, v28)
					if v32 != 0 {
						__phi28 = v32
						__phi29 = v28
						v28 = __phi28
						v29 = __phi29
						continue
					} else {
						break
					}
					break
				}
				v34 = v28
				v38 = v34 >> (uint(int32(31)) % 32)
				return v34 ^ v38 - v38
			} else {
				return int32(1)
			}
		}
	} else {
		if v16 != 0 {
			__phi28 = v16
			__phi29 = v17
			v28 = __phi28
			v29 = __phi29
			for {
				v32 = base.I32_rem_s(v29, v28)
				if v32 != 0 {
					__phi28 = v32
					__phi29 = v28
					v28 = __phi28
					v29 = __phi29
					continue
				} else {
					break
				}
				break
			}
			v34 = v28
		} else {
			v34 = v17
		}
		v38 = v34 >> (uint(int32(31)) % 32)
		return v34 ^ v38 - v38
	}
}
func F_int4inc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = v3 + int32(1)
	if v5 < v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int4inc_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4inc_1), int32(909), int32(_a_F_int4inc_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v5
	}
}
func F_int4out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_palloc(m, int32(12))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v2 {
			v17 = v2
			v18 = int32(0)
		} else {
			v12 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v12)
			v17 = int32(0) - v2
			v18 = int32(1)
		}
		v20 = F_pg_ultoa_n(m, v17, v4+v18)
		mBase = m.M
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4+(v20+v18)))) = uint8(v23)
		return v4
	}
}
func F_int4pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int4pl_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4pl_1), int32(829), int32(_a_F_int4pl_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v7
	}
}
func F_int4recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pq_getmsgint(m, v2, int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_int82ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v4 <= v3)
}
func F_int82lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 < v4)
}
func F_int82mul(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13862(m, l0, int32(_a_F_int82mul_0), int32(1069), int32(_a_F_int82mul_1), int32(_a_F_int82mul_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_int84eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 == v4)
}
func F_int8dec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == int32(0) {
		v36 = int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		switch v11 - int32(429) {
		case 0:
			v36 = int32(1)
		case 1:
			v36 = int32(2)
		default:
			v36 = int32(0)
		}
	}
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v40 = v38 - int64(1)
	v41 = base.B2i32(v38 <= v40)
	if v36 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40
		if v41 == int32(0) {
			v65 = v37
			return v65
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int8dec_0), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8dec_1), int32(774), int32(_a_F_int8dec_2))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		if v38 <= v40 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int8dec_0), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8dec_1), int32(787), int32(_a_F_int8dec_2))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v63 = F_Int64GetDatum(m, v40)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = v63
				return v65
			}
		}
	}
}
func F_int8inc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == int32(0) {
		v36 = int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		switch v11 - int32(429) {
		case 0:
			v36 = int32(1)
		case 1:
			v36 = int32(2)
		default:
			v36 = int32(0)
		}
	}
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v40 = v38 + int64(1)
	v41 = base.B2i32(v40 < v38)
	if v36 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40
		if v41 == int32(0) {
			v65 = v37
			return v65
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int8inc_0), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8inc_1), int32(736), int32(_a_F_int8inc_2))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		if v40 < v38 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int8inc_0), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8inc_1), int32(750), int32(_a_F_int8inc_2))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v63 = F_Int64GetDatum(m, v40)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = v63
				return v65
			}
		}
	}
}
func F_int8lcm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v39 int64
	_ = v39
	var __phi39 int64
	_ = __phi39
	var v40 int64
	_ = v40
	var __phi40 int64
	_ = __phi40
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v86 int64
	_ = v86
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v2 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v11 == v2 {
		v109 = v2
		v113 = F_Int64GetDatum(m, v109)
		mBase = m.M
		v116 = m.ExcPending
		if v116 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v113
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		if v15 == int64(0) {
			v109 = v2
			v113 = F_Int64GetDatum(m, v109)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v113
			}
		} else {
			v18 = int64(63)
			v19 = v11 >> (uint(v18) % 64)
			v23 = v15 >> (uint(v18) % 64)
			v26 = base.B2i32(base.Ui64(v23-(v23^v15)) < base.Ui64(v19-(v19^v11)))
			if base.Ui64(v23-(v23^v15)) < base.Ui64(v19-(v19^v11)) {
				v27 = v11
			} else {
				v27 = v15
			}
			if base.Ui64(v23-(v23^v15)) < base.Ui64(v19-(v19^v11)) {
				v28 = v15
			} else {
				v28 = v11
			}
			if v28 != int64(-9223372036854775807-1) {
				__phi39 = v27
				__phi40 = v28
				v39 = __phi39
				v40 = __phi40
				for {
					v43 = base.I64_rem_s(v40, v39)
					if v43 != int64(0) {
						__phi39 = v43
						__phi40 = v39
						v39 = __phi39
						v40 = __phi40
						continue
					} else {
						break
					}
					break
				}
				v47 = v39 >> (uint(int64(63)) % 64)
				v55 = v39 ^ v47 - v47
				v56 = base.I64_div_s(v11, v55)
				v57 = int64(63)
				v65 = int64(32)
				v66 = int64(base.Ui64(v15) >> (uint(v65) % 64))
				v68 = int64(base.Ui64(v56) >> (uint(v65) % 64))
				v71 = int64(4294967295)
				v72 = v15 & v71
				v74 = v56 & v71
				v75 = v72 * v74
				v79 = int64(base.Ui64(v75)>>(uint(v65)%64)) + v72*v68
				v86 = v74*v66 + v79&v71
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v56*(v15>>(uint(v57)%64)) + v56>>(uint(v57)%64)*v15 + v66*v68 + int64(base.Ui64(v79)>>(uint(v65)%64)) + int64(base.Ui64(v86)>>(uint(v65)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v8))) = v75&v71 | v86<<(uint(v65)%64)
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
				v98 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				if v97 != v98>>(uint(int64(63))%64) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int8lcm_1), int32(704), int32(_a_F_int8lcm_2))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if v98 == int64(-9223372036854775807-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v159 = m.ExcPending
							if v159 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_int8lcm_1), int32(710), int32(_a_F_int8lcm_2))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v105 = v98 >> (uint(int64(63)) % 64)
						v109 = v98 ^ v105 - v105
						v113 = F_Int64GetDatum(m, v109)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v113
						}
					}
				}
			} else {
				if v27&int64(9223372036854775807) == int64(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int8lcm_1), int32(636), int32(_a_F_int8lcm_3))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if v27 != int64(-1) {
						__phi39 = v27
						__phi40 = v28
						v39 = __phi39
						v40 = __phi40
						for {
							v43 = base.I64_rem_s(v40, v39)
							if v43 != int64(0) {
								__phi39 = v43
								__phi40 = v39
								v39 = __phi39
								v40 = __phi40
								continue
							} else {
								break
							}
							break
						}
						v47 = v39 >> (uint(int64(63)) % 64)
						v55 = v39 ^ v47 - v47
					} else {
						v55 = int64(1)
					}
					v56 = base.I64_div_s(v11, v55)
					v57 = int64(63)
					v65 = int64(32)
					v66 = int64(base.Ui64(v15) >> (uint(v65) % 64))
					v68 = int64(base.Ui64(v56) >> (uint(v65) % 64))
					v71 = int64(4294967295)
					v72 = v15 & v71
					v74 = v56 & v71
					v75 = v72 * v74
					v79 = int64(base.Ui64(v75)>>(uint(v65)%64)) + v72*v68
					v86 = v74*v66 + v79&v71
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v56*(v15>>(uint(v57)%64)) + v56>>(uint(v57)%64)*v15 + v66*v68 + int64(base.Ui64(v79)>>(uint(v65)%64)) + int64(base.Ui64(v86)>>(uint(v65)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v75&v71 | v86<<(uint(v65)%64)
					v97 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					v98 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
					if v97 != v98>>(uint(int64(63))%64) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_int8lcm_1), int32(704), int32(_a_F_int8lcm_2))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if v98 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_int8lcm_1), int32(710), int32(_a_F_int8lcm_2))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v105 = v98 >> (uint(int64(63)) % 64)
							v109 = v98 ^ v105 - v105
							v113 = F_Int64GetDatum(m, v109)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return v113
							}
						}
					}
				}
			}
		}
	}
}
func F_int8mi(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13861(m, l0, int32(_a_F_int8mi_0), int32(485), int32(_a_F_int8mi_1), int32(_a_F_int8mi_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_int8mod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = int64(1)
	v7 = v5 + v6
	if base.Ui64(v7) <= base.Ui64(v6) {
		if base.I32_wrap_i64(v7)-int32(1) != 0 {
			v35 = int64(0)
			v36 = F_Int64GetDatum(m, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				return v36
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int8mod_0), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8mod_1), int32(572), int32(_a_F_int8mod_2))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
		v34 = base.I64_rem_s(v33, v5)
		v35 = v34
		v36 = F_Int64GetDatum(m, v35)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			return v36
		}
	}
}
func F_int8pl(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13865(m, l0, int32(_a_F_int8pl_0), int32(471), int32(_a_F_int8pl_1), int32(_a_F_int8pl_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_intarray_match_first(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L21
	}
L2:
	;
	v7 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_ArrayGetNItemsSafe(m, v11, l0+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v7 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = (v19<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L11
L10:
	;
	v26 = v16
	goto L11
L11:
	;
	if int32(0) < v14 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	return int32(0)
L15:
	;
	v37 = v31 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0+v26+v31<<(uint(int32(2))%32))))
	if l1 == v41 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	return v37
L18:
	;
	goto L19
L19:
	;
	if v37 != v14 {
		v31 = v37
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_intarray_match_first_0), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_intarray_match_first_1), int32(344), int32(_a_F_intarray_match_first_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intarray_push_elem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_intarray_add_elem(m, v5, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v12 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					return v10
				}
			} else {
				return v10
			}
		}
	}
}
func F_interpret_function_parameter_list(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v473 int32
	_ = v473
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v931 int32
	_ = v931
	v14 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(80)
	m.G0 = v40
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v43 = v42
	goto L3
L2:
	;
	v43 = v14
	goto L3
L3:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v44
	v49 = v43 << (uint(int32(2)) % 32)
	v50 = F_palloc(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v52 = F_palloc(m, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v54 = F_palloc(m, v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v56 = F_palloc0(m, v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v58 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v58
	if l1 == v58 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v774 = F_buildoidvector(m, v50, v761)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L206
	}
L10:
	;
	v755 = v14
	v757 = v14
	v761 = v14
	v773 = int32(0)
	goto L9
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v62 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v88 = v14
	v90 = v14
	v91 = v14
	v94 = v14
	v98 = v14
	v100 = v14
	goto L23
L13:
	;
	v755 = v473
	v757 = v276
	v761 = v232
	v773 = base.B2i32(int32(0) < v293)
	goto L9
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L201
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L4
	} else {
		goto L196
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L191
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L186
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L181
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L176
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L171
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L4
	} else {
		goto L165
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L159
	}
L23:
	;
	v107 = v98 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v114 = F_LookupTypeName(m, l0, v112, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L153
	}
L25:
	;
	if v114 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+22)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v119)+82)))
	if v121 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L24
L28:
	;
	v151 = F_typeTypeId(m, v114)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L39
	}
L29:
	;
	if base.B2i32(l2 != int32(14)) == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	if l3 == int32(1) {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v128 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v128 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v135 = F_TypeNameToString(m, v112)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v135
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_0), v40-int32(-64))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(259), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L28
L39:
	;
	F_ReleaseCatCache(m, v114)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_interpret_function_parameter_list[0]))
	v159 = F_object_aclcheck(m, int32(1247), v151, v157, int64(256))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v159 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_aclcheck_error_type(m, v159, v151)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+12)))
	if v163 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v215 = v111 - int32(111)
	switch v215 {
	case 0, 5:
		v231 = int32(0)
		v232 = v94
		goto L65
	default:
		goto L66
	}
L49:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	if l3 != int32(29) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_3), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L62
	}
L52:
	;
	if l3 != int32(1) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_4), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L59
	}
L55:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_5), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(284), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(289), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(294), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	if v111 == int32(100) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	if int32(0) < v91 {
		goto L20
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v94<<(uint(int32(2))%32)))) = v151
	v222 = int32(1)
	v224 = v94 + v222
	if l5 == int32(0) {
		v231 = v222
		v232 = v224
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v228 = F_lappend_oid(m, v227, v151)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v228
	v231 = v222
	v232 = v224
	goto L65
L70:
	;
	v236 = int32(105)
	goto L72
L71:
	;
	v236 = v111
	goto L72
L72:
	;
	v238 = v236 - int32(105)
	v239 = int32(0)
	if base.B2i32(v238 == v239)|base.B2i32(v238 == int32(13)) == v239 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if l3 == int32(29) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v276 = v90
	goto L75
L75:
	;
	if v236 != int32(118) {
		v293 = v91
		goto L88
	} else {
		goto L89
	}
L76:
	;
	v276 = v90 + int32(1)
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v270
	goto L76
L78:
	;
	if v91 <= int32(0) {
		v270 = int32(2249)
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v90 != 0 {
		goto L76
	} else {
		goto L87
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_6), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(326), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v270 = v151
	goto L77
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107+v52))) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v107+v54))) = base.I32_extend8_s(v236)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v299 == int32(0) {
		v473 = v88
		goto L93
	} else {
		goto L94
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v151
	v281 = v91 + int32(1)
	if base.B2i32(v151 == int32(_a_F_interpret_function_parameter_list_7))|base.B2i32(base.Ui32(v151-int32(2276)) < base.Ui32(int32(2))) != 0 {
		v293 = v281
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v289 = F_get_element_type(m, v151)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v289 == int32(0) {
		goto L19
	} else {
		goto L92
	}
L92:
	;
	v293 = v281
	goto L88
L93:
	;
	if l9 != 0 {
		goto L129
	} else {
		goto L130
	}
L94:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v302 == int32(0) {
		v473 = v88
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v305 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v450 = F_cstring_to_text(m, v299)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L128
	}
L97:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v312 = int32(0)
	goto L98
L98:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v308+v312<<(uint(int32(2))%32))))
	if v350 == v110 {
		goto L96
	} else {
		goto L100
	}
L99:
	;
	goto L96
L100:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	if v353 == int32(100) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v356 = int32(105)
	goto L103
L102:
	;
	v356 = v353
	goto L103
L103:
	;
	if v238 != int32(13) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v410 = v312 + int32(1)
	if v410 != v305 {
		v312 = v410
		goto L98
	} else {
		goto L127
	}
L105:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v373 == int32(0) {
		goto L104
	} else {
		goto L117
	}
L106:
	;
	switch v215 {
	case 0, 5:
		goto L104
	default:
		goto L105
	}
L107:
	;
	v360 = v238
	goto L109
L108:
	;
	v360 = int32(0)
	goto L109
L109:
	;
	if v360 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	switch v356 - int32(105) {
	case 0, 13:
		goto L106
	default:
		goto L105
	case 6, 11:
		goto L104
	}
L111:
	;
	goto L112
L112:
	;
	v366 = v356 - int32(105)
	if v366 != int32(13) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v370 = v366
	goto L115
L114:
	;
	v370 = int32(0)
	goto L115
L115:
	;
	if v370 != 0 {
		goto L105
	} else {
		goto L116
	}
L116:
	;
	goto L106
L117:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v376 == int32(0) {
		goto L104
	} else {
		goto L118
	}
L118:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if base.B2i32(v381 == int32(0))|base.B2i32(v381 != v384) != 0 {
		v402 = v381
		v403 = v384
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v402-v403 == int32(0) {
		goto L18
	} else {
		goto L126
	}
L120:
	;
	goto L119
L121:
	;
	v387 = v373
	v388 = v299
	goto L122
L122:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+1)))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
	if v392 == int32(0) {
		v402 = v392
		v403 = v391
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v402 = v392
	v403 = v391
	goto L120
L124:
	;
	v395 = int32(1)
	if v392 == v391 {
		v387 = v387 + v395
		v388 = v388 + v395
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	goto L104
L127:
	;
	goto L99
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+v107))) = v450
	v473 = int32(1)
	goto L93
L129:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v492 != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if v503 != 0 {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	v496 = v492
	goto L134
L133:
	;
	v494 = F_pstrdup(m, int32(_a_F_interpret_function_parameter_list_8))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L135
	}
L134:
	;
	v497 = F_makeString(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L136
	}
L135:
	;
	v496 = v494
	goto L134
L136:
	;
	v499 = F_lappend(m, v491, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v499
	goto L131
L138:
	;
	v527 = v98 + int32(1)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v528 <= v527 {
		goto L13
	} else {
		goto L152
	}
L139:
	;
	if v231 == int32(0) {
		goto L17
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v231&v100 != 0 {
		goto L15
	} else {
		goto L150
	}
L142:
	;
	v507 = F_transformExpr(m, l0, v503, int32(31))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v510 = F_coerce_to_specific_type(m, l0, v507, v151, int32(_a_F_interpret_function_parameter_list_9))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	F_assign_expr_collations(m, l0, v510)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v514 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	v515 = F_contain_var_clause(m, v510)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v515 != 0 {
		goto L16
	} else {
		goto L148
	}
L148:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v518 = F_lappend(m, v517, v510)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v518
	v525 = int32(1)
	goto L138
L150:
	;
	if v100&base.B2i32(l3 == int32(29)) != 0 {
		goto L14
	} else {
		goto L151
	}
L151:
	;
	v525 = v100
	goto L138
L152:
	;
	v88 = v473
	v90 = v276
	v91 = v293
	v94 = v232
	v98 = v527
	v100 = v525
	goto L23
L153:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v537 = F_TypeNameToString(m, v112)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v537
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_10), v40+int32(32))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(246), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v560 = F_TypeNameToString(m, v112)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v560
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_11), v40+int32(48))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(253), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v583 = F_TypeNameToString(m, v112)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v583
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_12), v40)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(270), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_13), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v608)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(305), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_14), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(352), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v642
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_15), v40+int32(16))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(399), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_16), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(417), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(_a_F_interpret_function_parameter_list_17))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_18), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v687)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(433), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_19), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(458), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_20), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(469), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v774
	v777 = int32(0)
	if base.B2i32(v773 == v777)&base.B2i32(v757 <= v777) == v777 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	if v755 != 0 {
		goto L214
	} else {
		goto L215
	}
L208:
	;
	v785 = F_construct_array_builtin(m, v52, v43, int32(26))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L4
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v796 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v796
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v796
	goto L207
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v785
	v789 = F_construct_array_builtin(m, v54, v43, int32(18))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v789
	if v757 < int32(2) {
		goto L207
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = int32(2249)
	goto L207
L214:
	;
	if int32(0) < v43 {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v931 = int32(0)
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v931
	m.G0 = v40 + int32(80)
	return
L217:
	;
	v805 = int32(0)
	goto L220
L218:
	;
	goto L219
L219:
	;
	v891 = F_construct_array_builtin(m, v56, v43, int32(25))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L4
	} else {
		goto L227
	}
L220:
	;
	v842 = v56 + v805<<(uint(int32(2))%32)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	if v843 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L219
L222:
	;
	v847 = F_cstring_to_text(m, int32(_a_F_interpret_function_parameter_list_8))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L4
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v851 = v805 + int32(1)
	if v851 != v43 {
		v805 = v851
		goto L220
	} else {
		goto L226
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v842))) = v847
	goto L224
L226:
	;
	goto L221
L227:
	;
	v931 = v891
	goto L216
}
func F_inv_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(2224)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v24&int32(2) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l2 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L18
	} else {
		goto L110
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L18
	} else {
		goto L106
	}
L5:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(4398046509057)) <= base.Ui64(v29+base.I64_extend_i32_u(l2)) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v397 = v4
	goto L7
L7:
	;
	m.G0 = v22 + int32(2224)
	return v397
L8:
	;
	v36 = base.I32_wrap_i64(int64(base.Ui64(v29) >> (uint(int64(11)) % 64)))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[1]))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = v38
	goto L11
L10:
	;
	v42 = int32(0)
	goto L11
L11:
	;
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v45 = int32(_a_F_inv_write_0)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[2]))
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[2])) = v49
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v76 = v38
	goto L14
L14:
	;
	v79 = v22 + int32(80)
	v80 = F_CatalogOpenIndexes(m, v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L18
	} else {
		goto L24
	}
L15:
	;
	v61 = v38
	v62 = v41
	goto L17
L16:
	;
	v54 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v62 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[0])) = v54
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[1]))
	v61 = v54
	v62 = v60
	goto L17
L20:
	;
	v68 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v73 = v61
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[2])) = v46
	v76 = v73
	goto L14
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[1])) = v68
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v73 = v72
	goto L22
L24:
	;
	v83 = v22 + int32(2128)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v83, int32(1), int32(3), int32(184), v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	F_ScanKeyInit(m, v22+int32(2176), int32(2), int32(4), int32(150), v36)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v98 = v22 + int32(84)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[1]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v109 = F_systable_beginscan_ordered(m, v104, v106, v107, int32(2), v83)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v119 = v4
	v121 = int32(0)
	v123 = int32(1)
	v124 = v36
	v125 = v4
	goto L28
L28:
	;
	if v123 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	F_systable_endscan_ordered(m, v109)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L18
	} else {
		goto L103
	}
L30:
	;
	F_pfree(m, v370)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L18
	} else {
		goto L101
	}
L31:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v300 = base.I32_wrap_i64(v299)
	v302 = v300 & int32(2047)
	if v302 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L18
	} else {
		goto L80
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L18
	} else {
		goto L77
	}
L34:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v147 != v124 {
		v297 = v145
		v298 = v146
		goto L31
	} else {
		goto L43
	}
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+20)))
	if v139&int32(1) != 0 {
		goto L33
	} else {
		goto L42
	}
L36:
	;
	v133 = F_systable_getnext_ordered(m, v109, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	v136 = v125
	goto L38
L38:
	;
	if v119 != 0 {
		v145 = v119
		v146 = v136
		goto L34
	} else {
		goto L41
	}
L39:
	;
	if v133 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v136 = int32(0)
	goto L38
L41:
	;
	v297 = int32(0)
	v298 = v136
	goto L31
L42:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+22)))
	v145 = v138 + v142
	v146 = v133
	goto L34
L43:
	;
	v150 = v145 + int32(8)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+8)))
	v153 = v151 & int32(3)
	if v153 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v154 = F_detoast_attr(m, v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L18
	} else {
		goto L47
	}
L45:
	;
	v156 = v150
	goto L46
L46:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v159 = int32(base.Ui32(v157) >> (uint(int32(2)) % 32))
	v161 = v159 - int32(4)
	if base.Ui32(v157-int32(_a_F_inv_write_1)) <= base.Ui32(int32(-8197)) {
		goto L32
	} else {
		goto L48
	}
L47:
	;
	v156 = v154
	goto L46
L48:
	;
	if v161 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v79, v156+int32(4), v161)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if v153 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v156)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v174 = base.I32_wrap_i64(v171) & int32(2047)
	if v174 <= v161 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v212 = int32(2048) - v174
	v213 = l2 - v121
	if v212 < v213 {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v178 = v22 + int32(76) + v159
	v179 = int32(3)
	v181 = v174 - v161
	if v178&v179|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v181))|v181&v179 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v174+v79) <= base.Ui32(v178) {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	v202 = v181
	goto L60
L60:
	;
	if v202 == int32(0) {
		goto L56
	} else {
		goto L65
	}
L61:
	;
	v192 = v79 + v159
	v193 = v174 + v79
	if base.Ui32(v193) < base.Ui32(v192) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v195 = v192
	goto L64
L63:
	;
	v195 = v193
	goto L64
L64:
	;
	v202 = (v195+(v22+int32(76)^int32(-1))-v159)&int32(-4) + int32(4)
	goto L60
L65:
	;
	base.MemoryFill(m, v178, int32(0), v202)
	goto L56
L66:
	;
	v215 = v212
	goto L68
L67:
	;
	v215 = v213
	goto L68
L68:
	;
	if v215 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryCopy(m, v174+v79, l1+v121, v215)
	goto L71
L70:
	;
	goto L71
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v171 + base.I64_extend_i32_s(v215)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = int64(0)
	v224 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+60)) = uint16(v224)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+62)) = uint8(v224)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+56)) = uint16(v224)
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+58)) = uint8(v231)
	v234 = v215 + v174
	if v234 < v161 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v236 = v161
	goto L74
L73:
	;
	v236 = v234
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v236<<(uint(int32(2))%32) + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v22 + int32(76)
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	v254 = F_heap_modify_tuple(m, v146, v247, v22-int32(-64), v22+int32(60), v22+int32(56))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	F_CatalogTupleUpdateWithInfo(m, v257, v254+int32(4), v254, v80)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	v369 = v215
	v370 = v254
	v371 = v224
	v373 = v231
	v374 = int32(0)
	goto L30
L77:
	;
	F_errmsg_internal(m, int32(_a_F_inv_write_2), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_inv_write_3), int32(624), int32(_a_F_inv_write_4))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v283
	F_errmsg(m, int32(_a_F_inv_write_5), v22+int32(32))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L18
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_inv_write_3), int32(153), int32(_a_F_inv_write_6))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L18
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v329 = int32(2048) - v302
	v330 = l2 - v121
	if v329 < v330 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	if v300&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v302)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v312 = v302 + v79
	if base.Ui32(v98) < base.Ui32(v312) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v322 = v302
	goto L88
L88:
	;
	if v322 == int32(0) {
		goto L84
	} else {
		goto L92
	}
L89:
	;
	v314 = v312
	goto L91
L90:
	;
	v314 = v98
	goto L91
L91:
	;
	v322 = (v314-v22-int32(81))&int32(-4) + int32(4)
	goto L88
L92:
	;
	base.MemoryFill(m, v79, int32(0), v322)
	goto L84
L93:
	;
	v332 = v329
	goto L95
L94:
	;
	v332 = v330
	goto L95
L95:
	;
	if v332 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	base.MemoryCopy(m, v302+v79, l1+v121, v332)
	goto L98
L97:
	;
	goto L98
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v299 + base.I64_extend_i32_s(v332)
	v339 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+60)) = uint16(v339)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = (v332+v302)<<(uint(int32(2))%32) + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+62)) = uint8(v339)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v22 + int32(76)
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+52))
	v363 = F_heap_form_tuple(m, v358, v22-int32(-64), v22+int32(60))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L18
	} else {
		goto L99
	}
L99:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	F_CatalogTupleInsertWithInfo(m, v366, v363, v80)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L18
	} else {
		goto L100
	}
L100:
	;
	v369 = v332
	v370 = v363
	v371 = v297
	v373 = v339
	v374 = v298
	goto L30
L101:
	;
	v381 = v369 + v121
	if v381 < l2 {
		v119 = v371
		v121 = v381
		v123 = v373
		v124 = v124 + int32(1)
		v125 = v374
		goto L28
	} else {
		goto L102
	}
L102:
	;
	goto L29
L103:
	;
	F_CatalogCloseIndexes(m, v80)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L18
	} else {
		goto L104
	}
L104:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L18
	} else {
		goto L105
	}
L105:
	;
	v397 = v381
	goto L7
L106:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L18
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l2
	F_errmsg(m, int32(_a_F_inv_write_7), v22+int32(16))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L18
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_inv_write_3), int32(590), int32(_a_F_inv_write_4))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L18
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L18
	} else {
		goto L111
	}
L111:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v437
	F_errmsg(m, int32(_a_F_inv_write_8), v22)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_inv_write_3), int32(580), int32(_a_F_inv_write_4))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L18
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_is_admin_of_role(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_superuser_arg(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v23 = int32(1)
			m.G0 = v6 + int32(16)
			return v23
		} else {
			if l0 == l1 {
				v23 = int32(0)
				m.G0 = v6 + int32(16)
				return v23
			} else {
				v18 = F_roles_is_member_of(m, l0, int32(0), l1, v6+int32(12))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					v23 = base.B2i32(v20 != int32(0))
					m.G0 = v6 + int32(16)
					return v23
				}
			}
		}
	}
}
func F_is_code_in_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_is_code_in_table[0]))
	if base.Ui32(l0) < base.Ui32(v10) {
		v27 = v2
		m.G0 = v6 + int32(16)
		return v27
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_is_code_in_table[1]))
		if base.Ui32(v13) < base.Ui32(l0) {
			v27 = v2
			m.G0 = v6 + int32(16)
			return v27
		} else {
			v21 = F_bsearch(m, v6+int32(12), int32(_a_F_is_code_in_table_0), int32(34), int32(8), int32(_a_F_is_code_in_table_1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v27 = base.B2i32(v21 != int32(0))
				m.G0 = v6 + int32(16)
				return v27
			}
		}
	}
}
func F_is_dummy_rel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4 == v2 {
		v25 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		v8 = v7
		for {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui32(int32(2)) <= base.Ui32(v12-int32(301)) {
				break
			} else {
				v8 = v11 + int32(72)
				continue
			}
			break
		}
		if v12 != int32(290) {
			v25 = v2
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
			if v19 != 0 {
				v25 = v2
			} else {
				v25 = int32(1)
			}
		}
	}
	return v25
}
func F_is_leap(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	if int32(2147481747) < l0 {
		v6 = l0 - int32(2000)
	} else {
		v6 = l0
	}
	if v6&int32(3) != 0 {
		return int32(0)
	} else {
		v12 = v6 + int32(1900)
		v14 = base.I32_rem_s(v12, int32(100))
		if v14 != 0 {
			return int32(1)
		} else {
			v18 = base.I32_rem_s(v12, int32(400))
			return base.B2i32(v18 == int32(0))
		}
	}
}
func F_ismn_in(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13871(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_iso8859_1_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(8), int32(6))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v66)
	return v61 - v10
L4:
	;
	v61 = v10
	v63 = v9
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = v9
	v24 = v10
	v25 = v13
	goto L7
L7:
	;
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v24))))
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v61 = v54
	v63 = v51
	goto L3
L9:
	;
	if v8 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if int32(0) <= v29 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v61 = v24
	v63 = v23
	goto L3
L13:
	;
	goto L14
L14:
	;
	F_report_invalid_encoding(m, int32(8), v24, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v50 = v29
	v51 = v23 + int32(1)
	goto L18
L17:
	;
	v40 = v29 & int32(191)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v40)
	v50 = int32(base.Ui32(v29&int32(192))>>(uint(int32(6))%32)) | int32(-64)
	v51 = v23 + int32(2)
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v50)
	v53 = int32(1)
	v54 = v24 + v53
	if v53 < v25 {
		v23 = v51
		v24 = v54
		v25 = v25 - v53
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
}
func F_issn_cast_from_ean13(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13928(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_iswalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_iswalpha_0)) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswalpha[0]))))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v10<<(uint(int32(5))%32))+uint32(_c_F_iswalpha[0]))))
		return int32(base.Ui32(v14)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		return base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_iswalpha_1)))
	}
}
func F_iswgraph(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v37 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v37 = base.B2i32(v30 != int32(0))
	goto L1
L6:
	;
	v10 = int32(_a_F_iswgraph_0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v20 = int32(_a_F_iswgraph_0)
	v21 = F_wcslen(m, v20)
	mBase = m.M
	v30 = v21<<(uint(int32(2))%32) + v20
	goto L5
L9:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v13 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v13 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if l0 != v13 {
		v10 = v10 + int32(4)
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L13
L15:
	;
	v19 = v10
	goto L17
L16:
	;
	v19 = int32(0)
	goto L17
L17:
	;
	v30 = v19
	goto L5
L18:
	;
	v75 = int32(0)
	goto L20
L19:
	;
	if base.Ui32(l0) <= base.Ui32(int32(254)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return v75
L21:
	;
	v75 = base.B2i32(v72 != int32(0))
	goto L20
L22:
	;
	v72 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v49 = int32(_a_F_iswgraph_1)
	v72 = base.B2i32(l0&v49 != v49)&base.B2i32(base.Ui32(l0-int32(_a_F_iswgraph_2)) < base.Ui32(int32(_a_F_iswgraph_3))) | (base.B2i32(base.Ui32(l0-int32(_a_F_iswgraph_4)) < base.Ui32(int32(_a_F_iswgraph_5))) | base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_iswgraph_6))) | base.B2i32(base.Ui32(l0-int32(_a_F_iswgraph_7)) < base.Ui32(int32(_a_F_iswgraph_8))))
	goto L21
}
func F_italian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v499 int32
	_ = v499
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v686 int32
	_ = v686
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v805 int32
	_ = v805
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v924 int32
	_ = v924
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v957 int32
	_ = v957
	var v968 int32
	_ = v968
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1045 int32
	_ = v1045
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1083 int32
	_ = v1083
	var v1090 int32
	_ = v1090
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1112 int32
	_ = v1112
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1163 int32
	_ = v1163
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1197 int32
	_ = v1197
	var v1208 int32
	_ = v1208
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1286 int32
	_ = v1286
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1323 int32
	_ = v1323
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1404 int32
	_ = v1404
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1441 int32
	_ = v1441
	var v1448 int32
	_ = v1448
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1472 int32
	_ = v1472
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1523 int32
	_ = v1523
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1549 int32
	_ = v1549
	var v1556 int32
	_ = v1556
	var v1567 int32
	_ = v1567
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1644 int32
	_ = v1644
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1670 int32
	_ = v1670
	var v1682 int32
	_ = v1682
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1779 int32
	_ = v1779
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1830 int32
	_ = v1830
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1856 int32
	_ = v1856
	var v1863 int32
	_ = v1863
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1901 int32
	_ = v1901
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1952 int32
	_ = v1952
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1978 int32
	_ = v1978
	var v1986 int32
	_ = v1986
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2026 int32
	_ = v2026
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2064 int32
	_ = v2064
	var v2077 int32
	_ = v2077
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2103 int32
	_ = v2103
	var v2110 int32
	_ = v2110
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2148 int32
	_ = v2148
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2199 int32
	_ = v2199
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2225 int32
	_ = v2225
	var v2233 int32
	_ = v2233
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2314 int32
	_ = v2314
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2534 int32
	_ = v2534
	var v2537 int32
	_ = v2537
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2569 int32
	_ = v2569
	var v2573 int32
	_ = v2573
	var v2576 int32
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2616 int32
	_ = v2616
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2702 int32
	_ = v2702
	var v2704 int32
	_ = v2704
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2724 int32
	_ = v2724
	var v2740 int32
	_ = v2740
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2787 int32
	_ = v2787
	var v2789 int32
	_ = v2789
	var v2791 int32
	_ = v2791
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2835 int32
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2850 int32
	_ = v2850
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2895 int32
	_ = v2895
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2907 int32
	_ = v2907
	var v2923 int32
	_ = v2923
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2940 int32
	_ = v2940
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2973 int32
	_ = v2973
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2991 int32
	_ = v2991
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3005 int32
	_ = v3005
	var v3007 int32
	_ = v3007
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3021 int32
	_ = v3021
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3043 int32
	_ = v3043
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3055 int32
	_ = v3055
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v8 = int32(6)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v12-v6 < v8 {
		v22 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v43 = v6
	goto L14
L2:
	;
	if v22 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	goto L2
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = F_memcmp(m, v16+v6, int32(_a_F_italian_UTF_8_stem_0), v8)
	mBase = m.M
	if v18 != 0 {
		v22 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + v6
	v22 = int32(1)
	goto L3
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 < v26 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v25
	v32 = F_slice_from_s(m, l0, int32(5), int32(_a_F_italian_UTF_8_stem_1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if int32(0) <= v32 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(1)
	goto L12
L11:
	;
	v38 = v32
	goto L12
L12:
	;
	return v38
L13:
	;
	return v3055
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
	v50 = F_find_among(m, l0, int32(_a_F_italian_UTF_8_stem_2), int32(7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v155 = v6
	goto L62
L16:
	;
	goto L15
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v52
	switch v50 - int32(1) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	case 5:
		goto L20
	case 6:
		goto L19
	default:
		goto L18
	}
L18:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = v150
	goto L14
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L40
L20:
	;
	v88 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L36
	}
L21:
	;
	v82 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_4))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L34
	}
L22:
	;
	v76 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_5))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L32
	}
L23:
	;
	v70 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L30
	}
L24:
	;
	v64 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_7))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L28
	}
L25:
	;
	v58 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_8))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if int32(0) <= v58 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v3055 = v58
	goto L13
L28:
	;
	if int32(0) <= v64 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v3055 = v64
	goto L13
L30:
	;
	if int32(0) <= v70 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	v3055 = v70
	goto L13
L32:
	;
	if int32(0) <= v76 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v3055 = v76
	goto L13
L34:
	;
	if int32(0) <= v82 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	v3055 = v82
	goto L13
L36:
	;
	if int32(0) <= v88 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v3055 = v88
	goto L13
L38:
	;
	if v145 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	goto L41
L41:
	;
	goto L42
L42:
	;
	v100 = v52
	v102 = int32(1)
	goto L45
L44:
	;
	v145 = v130
	goto L38
L45:
	;
	if v93 <= v100 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v145 = int32(-1)
	goto L38
L48:
	;
	goto L49
L49:
	;
	v107 = v100 + int32(1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v100))))
	if base.Ui32(v109) < base.Ui32(int32(192)) {
		v130 = v107
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v131 = int32(1)
	if v131 < v102 {
		v100 = v130
		v102 = v102 - v131
		goto L45
	} else {
		goto L57
	}
L51:
	;
	if v93 <= v107 {
		v130 = v107
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v116 = v107
	goto L53
L53:
	;
	v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92+v116))))
	if int32(-65) < v119 {
		v130 = v116
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v130 = v93
	goto L50
L55:
	;
	v123 = v116 + int32(1)
	if v123 != v93 {
		v116 = v123
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L46
L58:
	;
	goto L16
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v145
	goto L18
L61:
	;
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2616
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2616
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L652
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L67
L63:
	;
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2593
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(v2592)+8))
	if v2593 < v2595 {
		goto L61
	} else {
		goto L641
	}
L64:
	;
	goto L63
L65:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v274 != 0 {
		v549 = v275
		goto L90
	} else {
		goto L91
	}
L66:
	;
	v274 = v267
	goto L65
L67:
	;
	if v169 <= v155 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v267 = int32(0)
	goto L66
L69:
	;
	v274 = int32(-1)
	goto L65
L70:
	;
	goto L71
L71:
	;
	v185 = int32(1)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v170))))
	if base.Ui32(v187) < base.Ui32(int32(192)) {
		v244 = v187
		v245 = v185
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if int32(249) < v244 {
		v267 = v245
		goto L66
	} else {
		goto L85
	}
L73:
	;
	v191 = v155 + int32(1)
	if v191 == v169 {
		v244 = v187
		v245 = v185
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v170))))
	v196 = v194 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v187) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v170))))
	v212 = v210 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v187) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v200 = v155 + int32(2)
	if v200 != v169 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v244 = v187<<(uint(int32(6))%32)&int32(1984) | v196
	v245 = int32(2)
	goto L72
L79:
	;
	goto L78
L80:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v216))))
	v244 = v229&int32(63) | (v187<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v196<<(uint(int32(12))%32) | v212<<(uint(int32(6))%32))
	v245 = int32(4)
	goto L72
L81:
	;
	v216 = v155 + int32(3)
	if v216 != v169 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v244 = v187<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v196<<(uint(int32(6))%32) | v212
	v245 = int32(3)
	goto L72
L84:
	;
	goto L83
L85:
	;
	v249 = v244 - int32(97)
	if v249 < int32(0) {
		v267 = v245
		goto L66
	} else {
		goto L86
	}
L86:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v249)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v255)>>(uint(v249&int32(7))%32))&int32(1) == int32(0) {
		v267 = v245
		goto L66
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v245 + v155
	goto L88
L88:
	;
	goto L68
L89:
	;
	v2587 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_11))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L8
	} else {
		goto L639
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L155
L91:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v276
	if v276 == v275 {
		v415 = v275
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v276
	if v276 == v415 {
		goto L124
	} else {
		goto L125
	}
L93:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v276))))
	if v281 != int32(117) {
		v415 = v275
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v285 = v276 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v285
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L97
L95:
	;
	if v405 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L96:
	;
	v405 = v398
	goto L95
L97:
	;
	if v300 <= v285 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v398 = int32(0)
	goto L96
L99:
	;
	v405 = int32(-1)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v316 = int32(1)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v301))))
	if base.Ui32(v318) < base.Ui32(int32(192)) {
		v375 = v318
		v376 = v316
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if int32(249) < v375 {
		v398 = v376
		goto L96
	} else {
		goto L115
	}
L103:
	;
	v322 = v276 + int32(2)
	if v322 == v300 {
		v375 = v318
		v376 = v316
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v301))))
	v327 = v325 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v318) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v301))))
	v343 = v341 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v318) {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v331 = v276 + int32(3)
	if v331 != v300 {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v375 = v318<<(uint(int32(6))%32)&int32(1984) | v327
	v376 = int32(2)
	goto L102
L109:
	;
	goto L108
L110:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v347))))
	v375 = v360&int32(63) | (v318<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v327<<(uint(int32(12))%32) | v343<<(uint(int32(6))%32))
	v376 = int32(4)
	goto L102
L111:
	;
	v347 = v276 + int32(4)
	if v347 != v300 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v375 = v318<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v327<<(uint(int32(6))%32) | v343
	v376 = int32(3)
	goto L102
L114:
	;
	goto L113
L115:
	;
	v380 = v375 - int32(97)
	if v380 < int32(0) {
		v398 = v376
		goto L96
	} else {
		goto L116
	}
L116:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v380)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v386)>>(uint(v380&int32(7))%32))&int32(1) == int32(0) {
		v398 = v376
		goto L96
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v376 + v285
	goto L118
L118:
	;
	goto L98
L119:
	;
	v410 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_12))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L8
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v415 = v414
	goto L92
L122:
	;
	if v410 < int32(0) {
		v3055 = v410
		goto L13
	} else {
		goto L123
	}
L123:
	;
	goto L62
L124:
	;
	v549 = v276
	goto L90
L125:
	;
	goto L126
L126:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+v276))))
	if v420 != int32(105) {
		v549 = v415
		goto L90
	} else {
		goto L127
	}
L127:
	;
	v424 = v276 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v424
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L130
L128:
	;
	if v544 == int32(0) {
		goto L89
	} else {
		goto L152
	}
L129:
	;
	v544 = v537
	goto L128
L130:
	;
	if v439 <= v424 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v537 = int32(0)
	goto L129
L132:
	;
	v544 = int32(-1)
	goto L128
L133:
	;
	goto L134
L134:
	;
	v455 = int32(1)
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v440))))
	if base.Ui32(v457) < base.Ui32(int32(192)) {
		v514 = v457
		v515 = v455
		goto L135
	} else {
		goto L136
	}
L135:
	;
	if int32(249) < v514 {
		v537 = v515
		goto L129
	} else {
		goto L148
	}
L136:
	;
	v461 = v276 + int32(2)
	if v461 == v439 {
		v514 = v457
		v515 = v455
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+v440))))
	v466 = v464 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v457) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470+v440))))
	v482 = v480 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v457) {
		goto L144
	} else {
		goto L145
	}
L139:
	;
	v470 = v276 + int32(3)
	if v470 != v439 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v514 = v457<<(uint(int32(6))%32)&int32(1984) | v466
	v515 = int32(2)
	goto L135
L142:
	;
	goto L141
L143:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440+v486))))
	v514 = v499&int32(63) | (v457<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v466<<(uint(int32(12))%32) | v482<<(uint(int32(6))%32))
	v515 = int32(4)
	goto L135
L144:
	;
	v486 = v276 + int32(4)
	if v486 != v439 {
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v514 = v457<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v466<<(uint(int32(6))%32) | v482
	v515 = int32(3)
	goto L135
L147:
	;
	goto L146
L148:
	;
	v519 = v514 - int32(97)
	if v519 < int32(0) {
		v537 = v515
		goto L129
	} else {
		goto L149
	}
L149:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v519)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v525)>>(uint(v519&int32(7))%32))&int32(1) == int32(0) {
		v537 = v515
		goto L129
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v515 + v424
	goto L151
L151:
	;
	goto L131
L152:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v549 = v547
	goto L90
L153:
	;
	if int32(0) <= v603 {
		v155 = v603
		goto L62
	} else {
		goto L173
	}
L155:
	;
	goto L156
L156:
	;
	goto L157
L157:
	;
	v558 = v155
	v560 = int32(1)
	goto L160
L159:
	;
	v603 = v588
	goto L153
L160:
	;
	if v549 <= v558 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	goto L159
L162:
	;
	v603 = int32(-1)
	goto L153
L163:
	;
	goto L164
L164:
	;
	v565 = v558 + int32(1)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551+v558))))
	if base.Ui32(v567) < base.Ui32(int32(192)) {
		v588 = v565
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v589 = int32(1)
	if v589 < v560 {
		v558 = v588
		v560 = v560 - v589
		goto L160
	} else {
		goto L172
	}
L166:
	;
	if v549 <= v565 {
		v588 = v565
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v574 = v565
	goto L168
L168:
	;
	v577 = int32(*(*int8)(unsafe.Add(mBase, uint32(v551+v574))))
	if int32(-65) < v577 {
		v588 = v574
		goto L165
	} else {
		goto L170
	}
L169:
	;
	v588 = v549
	goto L165
L170:
	;
	v581 = v574 + int32(1)
	if v581 != v549 {
		v574 = v581
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	goto L161
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v607)+4)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v607)+8)) = v608
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v607))) = v611
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L180
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v613
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1779 = v613
	goto L436
L175:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1753)+8)) = v1752
	goto L174
L176:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1752 = v1749 + v1748
	goto L175
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v613
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L309
L178:
	;
	if v731 != 0 {
		goto L177
	} else {
		goto L202
	}
L179:
	;
	v731 = v724
	goto L178
L180:
	;
	if v626 <= v613 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v724 = int32(0)
	goto L179
L182:
	;
	v731 = int32(-1)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v642 = int32(1)
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613+v627))))
	if base.Ui32(v644) < base.Ui32(int32(192)) {
		v701 = v644
		v702 = v642
		goto L185
	} else {
		goto L186
	}
L185:
	;
	if int32(249) < v701 {
		v724 = v702
		goto L179
	} else {
		goto L198
	}
L186:
	;
	v648 = v613 + int32(1)
	if v648 == v626 {
		v701 = v644
		v702 = v642
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648+v627))))
	v653 = v651 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v644) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657+v627))))
	v669 = v667 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v644) {
		goto L194
	} else {
		goto L195
	}
L189:
	;
	v657 = v613 + int32(2)
	if v657 != v626 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v701 = v644<<(uint(int32(6))%32)&int32(1984) | v653
	v702 = int32(2)
	goto L185
L192:
	;
	goto L191
L193:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627+v673))))
	v701 = v686&int32(63) | (v644<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v653<<(uint(int32(12))%32) | v669<<(uint(int32(6))%32))
	v702 = int32(4)
	goto L185
L194:
	;
	v673 = v613 + int32(3)
	if v673 != v626 {
		goto L193
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v701 = v644<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v653<<(uint(int32(6))%32) | v669
	v702 = int32(3)
	goto L185
L197:
	;
	goto L196
L198:
	;
	v706 = v701 - int32(97)
	if v706 < int32(0) {
		v724 = v702
		goto L179
	} else {
		goto L199
	}
L199:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v706)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v712)>>(uint(v706&int32(7))%32))&int32(1) == int32(0) {
		v724 = v702
		goto L179
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v702 + v613
	goto L201
L201:
	;
	goto L181
L202:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L205
L203:
	;
	if v849 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L204:
	;
	v849 = v842
	goto L203
L205:
	;
	if v745 <= v732 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v842 = int32(0)
	goto L204
L207:
	;
	v849 = int32(-1)
	goto L203
L208:
	;
	goto L209
L209:
	;
	v761 = int32(1)
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732+v746))))
	if base.Ui32(v763) < base.Ui32(int32(192)) {
		v820 = v763
		v821 = v761
		goto L210
	} else {
		goto L211
	}
L210:
	;
	if int32(249) < v820 {
		goto L223
	} else {
		goto L224
	}
L211:
	;
	v767 = v732 + int32(1)
	if v767 == v745 {
		v820 = v763
		v821 = v761
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767+v746))))
	v772 = v770 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v763) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776+v746))))
	v788 = v786 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v763) {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	v776 = v732 + int32(2)
	if v776 != v745 {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v820 = v763<<(uint(int32(6))%32)&int32(1984) | v772
	v821 = int32(2)
	goto L210
L217:
	;
	goto L216
L218:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v792))))
	v820 = v805&int32(63) | (v763<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v772<<(uint(int32(12))%32) | v788<<(uint(int32(6))%32))
	v821 = int32(4)
	goto L210
L219:
	;
	v792 = v732 + int32(3)
	if v792 != v745 {
		goto L218
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v820 = v763<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v772<<(uint(int32(6))%32) | v788
	v821 = int32(3)
	goto L210
L222:
	;
	goto L221
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v821 + v732
	goto L227
L224:
	;
	v825 = v820 - int32(97)
	if v825 < int32(0) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v825)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v831)>>(uint(v825&int32(7))%32))&int32(1) != 0 {
		v842 = v821
		goto L204
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	goto L206
L228:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v873 = v863
	goto L233
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v732
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L259
L231:
	;
	if int32(0) <= v968 {
		v1748 = v968
		goto L176
	} else {
		goto L256
	}
L232:
	;
	v968 = v940
	goto L231
L233:
	;
	if v864 <= v873 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v968 = int32(-1)
	goto L231
L236:
	;
	goto L237
L237:
	;
	v880 = int32(1)
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873+v865))))
	if base.Ui32(v882) < base.Ui32(int32(192)) {
		v939 = v882
		v940 = v880
		goto L238
	} else {
		goto L239
	}
L238:
	;
	if int32(249) < v939 {
		goto L251
	} else {
		goto L252
	}
L239:
	;
	v886 = v873 + int32(1)
	if v886 == v864 {
		v939 = v882
		v940 = v880
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886+v865))))
	v891 = v889 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v882) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895+v865))))
	v907 = v905 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v882) {
		goto L247
	} else {
		goto L248
	}
L242:
	;
	v895 = v873 + int32(2)
	if v895 != v864 {
		goto L241
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v939 = v882<<(uint(int32(6))%32)&int32(1984) | v891
	v940 = int32(2)
	goto L238
L245:
	;
	goto L244
L246:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865+v911))))
	v939 = v924&int32(63) | (v882<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v891<<(uint(int32(12))%32) | v907<<(uint(int32(6))%32))
	v940 = int32(4)
	goto L238
L247:
	;
	v911 = v873 + int32(3)
	if v911 != v864 {
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v939 = v882<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v891<<(uint(int32(6))%32) | v907
	v940 = int32(3)
	goto L238
L250:
	;
	goto L249
L251:
	;
	v957 = v940 + v873
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v957
	v873 = v957
	goto L233
L252:
	;
	v944 = v939 - int32(97)
	if v944 < int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v944)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v950)>>(uint(v944&int32(7))%32))&int32(1) != 0 {
		goto L232
	} else {
		goto L254
	}
L254:
	;
	goto L251
L256:
	;
	goto L230
L257:
	;
	if v1090 != 0 {
		goto L177
	} else {
		goto L281
	}
L258:
	;
	v1090 = v1083
	goto L257
L259:
	;
	if v985 <= v732 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1083 = int32(0)
	goto L258
L261:
	;
	v1090 = int32(-1)
	goto L257
L262:
	;
	goto L263
L263:
	;
	v1001 = int32(1)
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732+v986))))
	if base.Ui32(v1003) < base.Ui32(int32(192)) {
		v1060 = v1003
		v1061 = v1001
		goto L264
	} else {
		goto L265
	}
L264:
	;
	if int32(249) < v1060 {
		v1083 = v1061
		goto L258
	} else {
		goto L277
	}
L265:
	;
	v1007 = v732 + int32(1)
	if v1007 == v985 {
		v1060 = v1003
		v1061 = v1001
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007+v986))))
	v1012 = v1010 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1003) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1016+v986))))
	v1028 = v1026 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1003) {
		goto L273
	} else {
		goto L274
	}
L268:
	;
	v1016 = v732 + int32(2)
	if v1016 != v985 {
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1060 = v1003<<(uint(int32(6))%32)&int32(1984) | v1012
	v1061 = int32(2)
	goto L264
L271:
	;
	goto L270
L272:
	;
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986+v1032))))
	v1060 = v1045&int32(63) | (v1003<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1012<<(uint(int32(12))%32) | v1028<<(uint(int32(6))%32))
	v1061 = int32(4)
	goto L264
L273:
	;
	v1032 = v732 + int32(3)
	if v1032 != v985 {
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1060 = v1003<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1012<<(uint(int32(6))%32) | v1028
	v1061 = int32(3)
	goto L264
L276:
	;
	goto L275
L277:
	;
	v1065 = v1060 - int32(97)
	if v1065 < int32(0) {
		v1083 = v1061
		goto L258
	} else {
		goto L278
	}
L278:
	;
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1065)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1071)>>(uint(v1065&int32(7))%32))&int32(1) == int32(0) {
		v1083 = v1061
		goto L258
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1061 + v732
	goto L280
L280:
	;
	goto L260
L281:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1112 = v1102
	goto L284
L282:
	;
	if int32(0) <= v1208 {
		v1748 = v1208
		goto L176
	} else {
		goto L306
	}
L283:
	;
	v1208 = v1179
	goto L282
L284:
	;
	if v1103 <= v1112 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1208 = int32(-1)
	goto L282
L287:
	;
	goto L288
L288:
	;
	v1119 = int32(1)
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112+v1104))))
	if base.Ui32(v1121) < base.Ui32(int32(192)) {
		v1178 = v1121
		v1179 = v1119
		goto L289
	} else {
		goto L290
	}
L289:
	;
	if int32(249) < v1178 {
		goto L283
	} else {
		goto L302
	}
L290:
	;
	v1125 = v1112 + int32(1)
	if v1125 == v1103 {
		v1178 = v1121
		v1179 = v1119
		goto L289
	} else {
		goto L291
	}
L291:
	;
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125+v1104))))
	v1130 = v1128 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1121) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134+v1104))))
	v1146 = v1144 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1121) {
		goto L298
	} else {
		goto L299
	}
L293:
	;
	v1134 = v1112 + int32(2)
	if v1134 != v1103 {
		goto L292
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v1178 = v1121<<(uint(int32(6))%32)&int32(1984) | v1130
	v1179 = int32(2)
	goto L289
L296:
	;
	goto L295
L297:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104+v1150))))
	v1178 = v1163&int32(63) | (v1121<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1130<<(uint(int32(12))%32) | v1146<<(uint(int32(6))%32))
	v1179 = int32(4)
	goto L289
L298:
	;
	v1150 = v1112 + int32(3)
	if v1150 != v1103 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1178 = v1121<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1130<<(uint(int32(6))%32) | v1146
	v1179 = int32(3)
	goto L289
L301:
	;
	goto L300
L302:
	;
	v1183 = v1178 - int32(97)
	if v1183 < int32(0) {
		goto L283
	} else {
		goto L303
	}
L303:
	;
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1183)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1189)>>(uint(v1183&int32(7))%32))&int32(1) == int32(0) {
		goto L283
	} else {
		goto L304
	}
L304:
	;
	v1197 = v1179 + v1112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1197
	v1112 = v1197
	goto L284
L306:
	;
	goto L177
L307:
	;
	if v1330 != 0 {
		goto L174
	} else {
		goto L332
	}
L308:
	;
	v1330 = v1323
	goto L307
L309:
	;
	if v1226 <= v613 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v1323 = int32(0)
	goto L308
L311:
	;
	v1330 = int32(-1)
	goto L307
L312:
	;
	goto L313
L313:
	;
	v1242 = int32(1)
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613+v1227))))
	if base.Ui32(v1244) < base.Ui32(int32(192)) {
		v1301 = v1244
		v1302 = v1242
		goto L314
	} else {
		goto L315
	}
L314:
	;
	if int32(249) < v1301 {
		goto L327
	} else {
		goto L328
	}
L315:
	;
	v1248 = v613 + int32(1)
	if v1248 == v1226 {
		v1301 = v1244
		v1302 = v1242
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248+v1227))))
	v1253 = v1251 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1244) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257+v1227))))
	v1269 = v1267 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1244) {
		goto L323
	} else {
		goto L324
	}
L318:
	;
	v1257 = v613 + int32(2)
	if v1257 != v1226 {
		goto L317
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1301 = v1244<<(uint(int32(6))%32)&int32(1984) | v1253
	v1302 = int32(2)
	goto L314
L321:
	;
	goto L320
L322:
	;
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227+v1273))))
	v1301 = v1286&int32(63) | (v1244<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1253<<(uint(int32(12))%32) | v1269<<(uint(int32(6))%32))
	v1302 = int32(4)
	goto L314
L323:
	;
	v1273 = v613 + int32(3)
	if v1273 != v1226 {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1301 = v1244<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1253<<(uint(int32(6))%32) | v1269
	v1302 = int32(3)
	goto L314
L326:
	;
	goto L325
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1302 + v613
	goto L331
L328:
	;
	v1306 = v1301 - int32(97)
	if v1306 < int32(0) {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1306)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1312)>>(uint(v1306&int32(7))%32))&int32(1) != 0 {
		v1323 = v1302
		goto L308
	} else {
		goto L330
	}
L330:
	;
	goto L327
L331:
	;
	goto L310
L332:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L335
L333:
	;
	if v1448 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L334:
	;
	v1448 = v1441
	goto L333
L335:
	;
	if v1344 <= v1331 {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1441 = int32(0)
	goto L334
L337:
	;
	v1448 = int32(-1)
	goto L333
L338:
	;
	goto L339
L339:
	;
	v1360 = int32(1)
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331+v1345))))
	if base.Ui32(v1362) < base.Ui32(int32(192)) {
		v1419 = v1362
		v1420 = v1360
		goto L340
	} else {
		goto L341
	}
L340:
	;
	if int32(249) < v1419 {
		goto L353
	} else {
		goto L354
	}
L341:
	;
	v1366 = v1331 + int32(1)
	if v1366 == v1344 {
		v1419 = v1362
		v1420 = v1360
		goto L340
	} else {
		goto L342
	}
L342:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366+v1345))))
	v1371 = v1369 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1362) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375+v1345))))
	v1387 = v1385 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1362) {
		goto L349
	} else {
		goto L350
	}
L344:
	;
	v1375 = v1331 + int32(2)
	if v1375 != v1344 {
		goto L343
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1419 = v1362<<(uint(int32(6))%32)&int32(1984) | v1371
	v1420 = int32(2)
	goto L340
L347:
	;
	goto L346
L348:
	;
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345+v1391))))
	v1419 = v1404&int32(63) | (v1362<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1371<<(uint(int32(12))%32) | v1387<<(uint(int32(6))%32))
	v1420 = int32(4)
	goto L340
L349:
	;
	v1391 = v1331 + int32(3)
	if v1391 != v1344 {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1419 = v1362<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1371<<(uint(int32(6))%32) | v1387
	v1420 = int32(3)
	goto L340
L352:
	;
	goto L351
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1420 + v1331
	goto L357
L354:
	;
	v1424 = v1419 - int32(97)
	if v1424 < int32(0) {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1424)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1430)>>(uint(v1424&int32(7))%32))&int32(1) != 0 {
		v1441 = v1420
		goto L334
	} else {
		goto L356
	}
L356:
	;
	goto L353
L357:
	;
	goto L336
L358:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1472 = v1462
	goto L363
L359:
	;
	goto L360
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1331
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L389
L361:
	;
	if int32(0) <= v1567 {
		v1748 = v1567
		goto L176
	} else {
		goto L386
	}
L362:
	;
	v1567 = v1539
	goto L361
L363:
	;
	if v1463 <= v1472 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1567 = int32(-1)
	goto L361
L366:
	;
	goto L367
L367:
	;
	v1479 = int32(1)
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1472+v1464))))
	if base.Ui32(v1481) < base.Ui32(int32(192)) {
		v1538 = v1481
		v1539 = v1479
		goto L368
	} else {
		goto L369
	}
L368:
	;
	if int32(249) < v1538 {
		goto L381
	} else {
		goto L382
	}
L369:
	;
	v1485 = v1472 + int32(1)
	if v1485 == v1463 {
		v1538 = v1481
		v1539 = v1479
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1485+v1464))))
	v1490 = v1488 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1481) {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	v1504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1494+v1464))))
	v1506 = v1504 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1481) {
		goto L377
	} else {
		goto L378
	}
L372:
	;
	v1494 = v1472 + int32(2)
	if v1494 != v1463 {
		goto L371
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	v1538 = v1481<<(uint(int32(6))%32)&int32(1984) | v1490
	v1539 = int32(2)
	goto L368
L375:
	;
	goto L374
L376:
	;
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464+v1510))))
	v1538 = v1523&int32(63) | (v1481<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1490<<(uint(int32(12))%32) | v1506<<(uint(int32(6))%32))
	v1539 = int32(4)
	goto L368
L377:
	;
	v1510 = v1472 + int32(3)
	if v1510 != v1463 {
		goto L376
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v1538 = v1481<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1490<<(uint(int32(6))%32) | v1506
	v1539 = int32(3)
	goto L368
L380:
	;
	goto L379
L381:
	;
	v1556 = v1539 + v1472
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1556
	v1472 = v1556
	goto L363
L382:
	;
	v1543 = v1538 - int32(97)
	if v1543 < int32(0) {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1543)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1549)>>(uint(v1543&int32(7))%32))&int32(1) != 0 {
		goto L362
	} else {
		goto L384
	}
L384:
	;
	goto L381
L386:
	;
	goto L360
L387:
	;
	if v1689 != 0 {
		goto L174
	} else {
		goto L411
	}
L388:
	;
	v1689 = v1682
	goto L387
L389:
	;
	if v1584 <= v1331 {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v1682 = int32(0)
	goto L388
L391:
	;
	v1689 = int32(-1)
	goto L387
L392:
	;
	goto L393
L393:
	;
	v1600 = int32(1)
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331+v1585))))
	if base.Ui32(v1602) < base.Ui32(int32(192)) {
		v1659 = v1602
		v1660 = v1600
		goto L394
	} else {
		goto L395
	}
L394:
	;
	if int32(249) < v1659 {
		v1682 = v1660
		goto L388
	} else {
		goto L407
	}
L395:
	;
	v1606 = v1331 + int32(1)
	if v1606 == v1584 {
		v1659 = v1602
		v1660 = v1600
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1606+v1585))))
	v1611 = v1609 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1602) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	v1625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615+v1585))))
	v1627 = v1625 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1602) {
		goto L403
	} else {
		goto L404
	}
L398:
	;
	v1615 = v1331 + int32(2)
	if v1615 != v1584 {
		goto L397
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v1659 = v1602<<(uint(int32(6))%32)&int32(1984) | v1611
	v1660 = int32(2)
	goto L394
L401:
	;
	goto L400
L402:
	;
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585+v1631))))
	v1659 = v1644&int32(63) | (v1602<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1611<<(uint(int32(12))%32) | v1627<<(uint(int32(6))%32))
	v1660 = int32(4)
	goto L394
L403:
	;
	v1631 = v1331 + int32(3)
	if v1631 != v1584 {
		goto L402
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1659 = v1602<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1611<<(uint(int32(6))%32) | v1627
	v1660 = int32(3)
	goto L394
L406:
	;
	goto L405
L407:
	;
	v1664 = v1659 - int32(97)
	if v1664 < int32(0) {
		v1682 = v1660
		goto L388
	} else {
		goto L408
	}
L408:
	;
	v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1664)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1670)>>(uint(v1664&int32(7))%32))&int32(1) == int32(0) {
		v1682 = v1660
		goto L388
	} else {
		goto L409
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1660 + v1331
	goto L410
L410:
	;
	goto L390
L411:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L414
L412:
	;
	if int32(0) <= v1744 {
		v1752 = v1744
		goto L175
	} else {
		goto L432
	}
L414:
	;
	goto L415
L415:
	;
	goto L416
L416:
	;
	v1699 = v1691
	v1701 = int32(1)
	goto L419
L418:
	;
	v1744 = v1729
	goto L412
L419:
	;
	if v1692 <= v1699 {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	goto L418
L421:
	;
	v1744 = int32(-1)
	goto L412
L422:
	;
	goto L423
L423:
	;
	v1706 = v1699 + int32(1)
	v1708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1690+v1699))))
	if base.Ui32(v1708) < base.Ui32(int32(192)) {
		v1729 = v1706
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1730 = int32(1)
	if v1730 < v1701 {
		v1699 = v1729
		v1701 = v1701 - v1730
		goto L419
	} else {
		goto L431
	}
L425:
	;
	if v1692 <= v1706 {
		v1729 = v1706
		goto L424
	} else {
		goto L426
	}
L426:
	;
	v1715 = v1706
	goto L427
L427:
	;
	v1718 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1690+v1715))))
	if int32(-65) < v1718 {
		v1729 = v1715
		goto L424
	} else {
		goto L429
	}
L428:
	;
	v1729 = v1692
	goto L424
L429:
	;
	v1722 = v1715 + int32(1)
	if v1722 != v1692 {
		v1715 = v1722
		goto L427
	} else {
		goto L430
	}
L430:
	;
	goto L428
L431:
	;
	goto L420
L432:
	;
	goto L174
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v613
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2253
	v2257 = v2253 - int32(1)
	if v2257 <= v613 {
		goto L536
	} else {
		goto L537
	}
L434:
	;
	if v1874 < int32(0) {
		goto L433
	} else {
		goto L459
	}
L435:
	;
	v1874 = v1846
	goto L434
L436:
	;
	if v1770 <= v1779 {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v1874 = int32(-1)
	goto L434
L439:
	;
	goto L440
L440:
	;
	v1786 = int32(1)
	v1788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1779+v1771))))
	if base.Ui32(v1788) < base.Ui32(int32(192)) {
		v1845 = v1788
		v1846 = v1786
		goto L441
	} else {
		goto L442
	}
L441:
	;
	if int32(249) < v1845 {
		goto L454
	} else {
		goto L455
	}
L442:
	;
	v1792 = v1779 + int32(1)
	if v1792 == v1770 {
		v1845 = v1788
		v1846 = v1786
		goto L441
	} else {
		goto L443
	}
L443:
	;
	v1795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1792+v1771))))
	v1797 = v1795 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1788) {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	v1811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1801+v1771))))
	v1813 = v1811 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1788) {
		goto L450
	} else {
		goto L451
	}
L445:
	;
	v1801 = v1779 + int32(2)
	if v1801 != v1770 {
		goto L444
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	v1845 = v1788<<(uint(int32(6))%32)&int32(1984) | v1797
	v1846 = int32(2)
	goto L441
L448:
	;
	goto L447
L449:
	;
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1771+v1817))))
	v1845 = v1830&int32(63) | (v1788<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1797<<(uint(int32(12))%32) | v1813<<(uint(int32(6))%32))
	v1846 = int32(4)
	goto L441
L450:
	;
	v1817 = v1779 + int32(3)
	if v1817 != v1770 {
		goto L449
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	v1845 = v1788<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1797<<(uint(int32(6))%32) | v1813
	v1846 = int32(3)
	goto L441
L453:
	;
	goto L452
L454:
	;
	v1863 = v1846 + v1779
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1863
	v1779 = v1863
	goto L436
L455:
	;
	v1850 = v1845 - int32(97)
	if v1850 < int32(0) {
		goto L454
	} else {
		goto L456
	}
L456:
	;
	v1856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1850)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1856)>>(uint(v1850&int32(7))%32))&int32(1) != 0 {
		goto L435
	} else {
		goto L457
	}
L457:
	;
	goto L454
L459:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1878 = v1877 + v1874
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1878
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1901 = v1878
	goto L462
L460:
	;
	if v1997 < int32(0) {
		goto L433
	} else {
		goto L484
	}
L461:
	;
	v1997 = v1968
	goto L460
L462:
	;
	if v1892 <= v1901 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1997 = int32(-1)
	goto L460
L465:
	;
	goto L466
L466:
	;
	v1908 = int32(1)
	v1910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1901+v1893))))
	if base.Ui32(v1910) < base.Ui32(int32(192)) {
		v1967 = v1910
		v1968 = v1908
		goto L467
	} else {
		goto L468
	}
L467:
	;
	if int32(249) < v1967 {
		goto L461
	} else {
		goto L480
	}
L468:
	;
	v1914 = v1901 + int32(1)
	if v1914 == v1892 {
		v1967 = v1910
		v1968 = v1908
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v1917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1914+v1893))))
	v1919 = v1917 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1910) {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	v1933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1923+v1893))))
	v1935 = v1933 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1910) {
		goto L476
	} else {
		goto L477
	}
L471:
	;
	v1923 = v1901 + int32(2)
	if v1923 != v1892 {
		goto L470
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	v1967 = v1910<<(uint(int32(6))%32)&int32(1984) | v1919
	v1968 = int32(2)
	goto L467
L474:
	;
	goto L473
L475:
	;
	v1952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1893+v1939))))
	v1967 = v1952&int32(63) | (v1910<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v1919<<(uint(int32(12))%32) | v1935<<(uint(int32(6))%32))
	v1968 = int32(4)
	goto L467
L476:
	;
	v1939 = v1901 + int32(3)
	if v1939 != v1892 {
		goto L475
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v1967 = v1910<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v1919<<(uint(int32(6))%32) | v1935
	v1968 = int32(3)
	goto L467
L479:
	;
	goto L478
L480:
	;
	v1972 = v1967 - int32(97)
	if v1972 < int32(0) {
		goto L461
	} else {
		goto L481
	}
L481:
	;
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1972)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1978)>>(uint(v1972&int32(7))%32))&int32(1) == int32(0) {
		goto L461
	} else {
		goto L482
	}
L482:
	;
	v1986 = v1968 + v1901
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1986
	v1901 = v1986
	goto L462
L484:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2001 = v2000 + v1997
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2001
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+4)) = v2001
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2026 = v2016
	goto L487
L485:
	;
	if v2121 < int32(0) {
		goto L433
	} else {
		goto L510
	}
L486:
	;
	v2121 = v2093
	goto L485
L487:
	;
	if v2017 <= v2026 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2121 = int32(-1)
	goto L485
L490:
	;
	goto L491
L491:
	;
	v2033 = int32(1)
	v2035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026+v2018))))
	if base.Ui32(v2035) < base.Ui32(int32(192)) {
		v2092 = v2035
		v2093 = v2033
		goto L492
	} else {
		goto L493
	}
L492:
	;
	if int32(249) < v2092 {
		goto L505
	} else {
		goto L506
	}
L493:
	;
	v2039 = v2026 + int32(1)
	if v2039 == v2017 {
		v2092 = v2035
		v2093 = v2033
		goto L492
	} else {
		goto L494
	}
L494:
	;
	v2042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2039+v2018))))
	v2044 = v2042 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2035) {
		goto L496
	} else {
		goto L497
	}
L495:
	;
	v2058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048+v2018))))
	v2060 = v2058 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2035) {
		goto L501
	} else {
		goto L502
	}
L496:
	;
	v2048 = v2026 + int32(2)
	if v2048 != v2017 {
		goto L495
	} else {
		goto L499
	}
L497:
	;
	goto L498
L498:
	;
	v2092 = v2035<<(uint(int32(6))%32)&int32(1984) | v2044
	v2093 = int32(2)
	goto L492
L499:
	;
	goto L498
L500:
	;
	v2077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2018+v2064))))
	v2092 = v2077&int32(63) | (v2035<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v2044<<(uint(int32(12))%32) | v2060<<(uint(int32(6))%32))
	v2093 = int32(4)
	goto L492
L501:
	;
	v2064 = v2026 + int32(3)
	if v2064 != v2017 {
		goto L500
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v2092 = v2035<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v2044<<(uint(int32(6))%32) | v2060
	v2093 = int32(3)
	goto L492
L504:
	;
	goto L503
L505:
	;
	v2110 = v2093 + v2026
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2110
	v2026 = v2110
	goto L487
L506:
	;
	v2097 = v2092 - int32(97)
	if v2097 < int32(0) {
		goto L505
	} else {
		goto L507
	}
L507:
	;
	v2103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2097)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2103)>>(uint(v2097&int32(7))%32))&int32(1) != 0 {
		goto L486
	} else {
		goto L508
	}
L508:
	;
	goto L505
L510:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2125 = v2124 + v2121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2125
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2148 = v2125
	goto L513
L511:
	;
	if v2244 < int32(0) {
		goto L433
	} else {
		goto L535
	}
L512:
	;
	v2244 = v2215
	goto L511
L513:
	;
	if v2139 <= v2148 {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	v2244 = int32(-1)
	goto L511
L516:
	;
	goto L517
L517:
	;
	v2155 = int32(1)
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2148+v2140))))
	if base.Ui32(v2157) < base.Ui32(int32(192)) {
		v2214 = v2157
		v2215 = v2155
		goto L518
	} else {
		goto L519
	}
L518:
	;
	if int32(249) < v2214 {
		goto L512
	} else {
		goto L531
	}
L519:
	;
	v2161 = v2148 + int32(1)
	if v2161 == v2139 {
		v2214 = v2157
		v2215 = v2155
		goto L518
	} else {
		goto L520
	}
L520:
	;
	v2164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161+v2140))))
	v2166 = v2164 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2157) {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	v2180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170+v2140))))
	v2182 = v2180 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2157) {
		goto L527
	} else {
		goto L528
	}
L522:
	;
	v2170 = v2148 + int32(2)
	if v2170 != v2139 {
		goto L521
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	v2214 = v2157<<(uint(int32(6))%32)&int32(1984) | v2166
	v2215 = int32(2)
	goto L518
L525:
	;
	goto L524
L526:
	;
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140+v2186))))
	v2214 = v2199&int32(63) | (v2157<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_9) | v2166<<(uint(int32(12))%32) | v2182<<(uint(int32(6))%32))
	v2215 = int32(4)
	goto L518
L527:
	;
	v2186 = v2148 + int32(3)
	if v2186 != v2139 {
		goto L526
	} else {
		goto L530
	}
L528:
	;
	goto L529
L529:
	;
	v2214 = v2157<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v2166<<(uint(int32(6))%32) | v2182
	v2215 = int32(3)
	goto L518
L530:
	;
	goto L529
L531:
	;
	v2219 = v2214 - int32(97)
	if v2219 < int32(0) {
		goto L512
	} else {
		goto L532
	}
L532:
	;
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2219)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2225)>>(uint(v2219&int32(7))%32))&int32(1) == int32(0) {
		goto L512
	} else {
		goto L533
	}
L533:
	;
	v2233 = v2215 + v2148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2233
	v2148 = v2233
	goto L513
L535:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2247))) = v2248 + v2244
	goto L433
L536:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2314
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2314
	v2319 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_13), int32(51))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L8
	} else {
		goto L552
	}
L537:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259+v2257))))
	if base.B2i32(v2261&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2261)%32)&int32(_a_F_italian_UTF_8_stem_14) == int32(0)) != 0 {
		goto L536
	} else {
		goto L538
	}
L538:
	;
	v2275 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_15), int32(37))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L8
	} else {
		goto L539
	}
L539:
	;
	if v2275 == int32(0) {
		goto L536
	} else {
		goto L540
	}
L540:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2279
	v2282 = v2279 - int32(1)
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2282 <= v2283 {
		goto L536
	} else {
		goto L541
	}
L541:
	;
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2285+v2282))))
	switch v2287 - int32(111) {
	case 0, 3:
		goto L542
	default:
		goto L536
	}
L542:
	;
	v2292 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_16), int32(5))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L8
	} else {
		goto L543
	}
L543:
	;
	if v2292 == int32(0) {
		goto L536
	} else {
		goto L544
	}
L544:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+8))
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2298 < v2297 {
		goto L536
	} else {
		goto L545
	}
L545:
	;
	switch v2292 - int32(1) {
	case 0:
		goto L547
	case 1:
		goto L546
	default:
		goto L536
	}
L546:
	;
	v2308 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_17))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L8
	} else {
		goto L550
	}
L547:
	;
	v2302 = F_slice_del(m, l0)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L8
	} else {
		goto L548
	}
L548:
	;
	if int32(0) <= v2302 {
		goto L536
	} else {
		goto L549
	}
L549:
	;
	v3055 = v2302
	goto L13
L550:
	;
	if v2308 < int32(0) {
		v3055 = v2308
		goto L13
	} else {
		goto L551
	}
L551:
	;
	goto L536
L552:
	;
	if v2319 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2592 = v2323
	goto L64
L554:
	;
	goto L555
L555:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2324
	switch v2319 - int32(1) {
	case 0:
		goto L564
	case 1:
		goto L563
	case 2:
		goto L562
	case 3:
		goto L561
	case 4:
		goto L560
	case 5:
		goto L559
	case 6:
		goto L558
	case 7:
		goto L557
	case 8:
		goto L556
	default:
		goto L61
	}
L556:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v2520)))
	if v2324 < v2521 {
		v2592 = v2520
		goto L64
	} else {
		goto L620
	}
L557:
	;
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2478)))
	if v2324 < v2479 {
		v2592 = v2478
		goto L64
	} else {
		goto L610
	}
L558:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2405)+4))
	if v2324 < v2406 {
		v2592 = v2405
		goto L64
	} else {
		goto L591
	}
L559:
	;
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v2398)+8))
	if v2324 < v2399 {
		v2592 = v2398
		goto L64
	} else {
		goto L588
	}
L560:
	;
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2389)))
	if v2324 < v2390 {
		v2592 = v2389
		goto L64
	} else {
		goto L585
	}
L561:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	if v2324 < v2381 {
		v2592 = v2380
		goto L64
	} else {
		goto L582
	}
L562:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v2371)))
	if v2324 < v2372 {
		v2592 = v2371
		goto L64
	} else {
		goto L579
	}
L563:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2335)))
	if v2324 < v2336 {
		v2592 = v2335
		goto L64
	} else {
		goto L568
	}
L564:
	;
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2328)))
	if v2324 < v2329 {
		v2592 = v2328
		goto L64
	} else {
		goto L565
	}
L565:
	;
	v2331 = F_slice_del(m, l0)
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L8
	} else {
		goto L566
	}
L566:
	;
	if int32(0) <= v2331 {
		goto L61
	} else {
		goto L567
	}
L567:
	;
	v3055 = v2331
	goto L13
L568:
	;
	v2338 = F_slice_del(m, l0)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L8
	} else {
		goto L569
	}
L569:
	;
	if v2338 < int32(0) {
		v3055 = v2338
		goto L13
	} else {
		goto L570
	}
L570:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2342
	v2344 = int32(2)
	v2346 = int32(0)
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2342-v2349 < v2344 {
		v2359 = v2346
		goto L572
	} else {
		goto L573
	}
L571:
	;
	if v2359 == int32(0) {
		goto L61
	} else {
		goto L575
	}
L572:
	;
	goto L571
L573:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2355 = F_memcmp(m, v2352+v2342-v2344, int32(_a_F_italian_UTF_8_stem_18), v2344)
	mBase = m.M
	if v2355 != 0 {
		v2359 = v2346
		goto L572
	} else {
		goto L574
	}
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2342 - v2344
	v2359 = int32(1)
	goto L572
L575:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2362
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v2364)))
	if v2362 < v2365 {
		goto L61
	} else {
		goto L576
	}
L576:
	;
	v2367 = F_slice_del(m, l0)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L8
	} else {
		goto L577
	}
L577:
	;
	if int32(0) <= v2367 {
		goto L61
	} else {
		goto L578
	}
L578:
	;
	v3055 = v2367
	goto L13
L579:
	;
	v2376 = F_slice_from_s(m, l0, int32(3), int32(_a_F_italian_UTF_8_stem_19))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L8
	} else {
		goto L580
	}
L580:
	;
	if int32(0) <= v2376 {
		goto L61
	} else {
		goto L581
	}
L581:
	;
	v3055 = v2376
	goto L13
L582:
	;
	v2385 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_20))
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L8
	} else {
		goto L583
	}
L583:
	;
	if int32(0) <= v2385 {
		goto L61
	} else {
		goto L584
	}
L584:
	;
	v3055 = v2385
	goto L13
L585:
	;
	v2394 = F_slice_from_s(m, l0, int32(4), int32(_a_F_italian_UTF_8_stem_21))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L8
	} else {
		goto L586
	}
L586:
	;
	if int32(0) <= v2394 {
		goto L61
	} else {
		goto L587
	}
L587:
	;
	v3055 = v2394
	goto L13
L588:
	;
	v2401 = F_slice_del(m, l0)
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L8
	} else {
		goto L589
	}
L589:
	;
	if int32(0) <= v2401 {
		goto L61
	} else {
		goto L590
	}
L590:
	;
	v3055 = v2401
	goto L13
L591:
	;
	v2408 = F_slice_del(m, l0)
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L8
	} else {
		goto L592
	}
L592:
	;
	if v2408 < int32(0) {
		v3055 = v2408
		goto L13
	} else {
		goto L593
	}
L593:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2412
	v2415 = v2412 - int32(1)
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2415 <= v2416 {
		goto L61
	} else {
		goto L594
	}
L594:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2418+v2415))))
	if base.B2i32(v2420&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2420)%32)&int32(_a_F_italian_UTF_8_stem_22) == int32(0)) != 0 {
		goto L61
	} else {
		goto L595
	}
L595:
	;
	v2434 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_23), int32(4))
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L8
	} else {
		goto L596
	}
L596:
	;
	if v2434 == int32(0) {
		goto L61
	} else {
		goto L597
	}
L597:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2438
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2440)))
	if v2438 < v2441 {
		goto L61
	} else {
		goto L598
	}
L598:
	;
	v2443 = F_slice_del(m, l0)
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L8
	} else {
		goto L599
	}
L599:
	;
	if v2443 < int32(0) {
		v3055 = v2443
		goto L13
	} else {
		goto L600
	}
L600:
	;
	if v2434 != int32(1) {
		goto L61
	} else {
		goto L601
	}
L601:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2449
	v2451 = int32(2)
	v2453 = int32(0)
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2449-v2456 < v2451 {
		v2466 = v2453
		goto L603
	} else {
		goto L604
	}
L602:
	;
	if v2466 == int32(0) {
		goto L61
	} else {
		goto L606
	}
L603:
	;
	goto L602
L604:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2462 = F_memcmp(m, v2459+v2449-v2451, int32(_a_F_italian_UTF_8_stem_24), v2451)
	mBase = m.M
	if v2462 != 0 {
		v2466 = v2453
		goto L603
	} else {
		goto L605
	}
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2449 - v2451
	v2466 = int32(1)
	goto L603
L606:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2469
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v2471)))
	if v2469 < v2472 {
		goto L61
	} else {
		goto L607
	}
L607:
	;
	v2474 = F_slice_del(m, l0)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L8
	} else {
		goto L608
	}
L608:
	;
	if int32(0) <= v2474 {
		goto L61
	} else {
		goto L609
	}
L609:
	;
	v3055 = v2474
	goto L13
L610:
	;
	v2481 = F_slice_del(m, l0)
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L8
	} else {
		goto L611
	}
L611:
	;
	if v2481 < int32(0) {
		v3055 = v2481
		goto L13
	} else {
		goto L612
	}
L612:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2485
	v2488 = v2485 - int32(1)
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2488 <= v2489 {
		goto L61
	} else {
		goto L613
	}
L613:
	;
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491+v2488))))
	if base.B2i32(v2493&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2493)%32)&int32(_a_F_italian_UTF_8_stem_25) == int32(0)) != 0 {
		goto L61
	} else {
		goto L614
	}
L614:
	;
	v2507 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_26), int32(3))
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L8
	} else {
		goto L615
	}
L615:
	;
	if v2507 == int32(0) {
		goto L61
	} else {
		goto L616
	}
L616:
	;
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2511
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2513)))
	if v2511 < v2514 {
		goto L61
	} else {
		goto L617
	}
L617:
	;
	v2516 = F_slice_del(m, l0)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L8
	} else {
		goto L618
	}
L618:
	;
	if int32(0) <= v2516 {
		goto L61
	} else {
		goto L619
	}
L619:
	;
	v3055 = v2516
	goto L13
L620:
	;
	v2523 = F_slice_del(m, l0)
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L8
	} else {
		goto L621
	}
L621:
	;
	if v2523 < int32(0) {
		v3055 = v2523
		goto L13
	} else {
		goto L622
	}
L622:
	;
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2527
	v2529 = int32(2)
	v2531 = int32(0)
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2527-v2534 < v2529 {
		v2544 = v2531
		goto L624
	} else {
		goto L625
	}
L623:
	;
	if v2544 == int32(0) {
		goto L61
	} else {
		goto L627
	}
L624:
	;
	goto L623
L625:
	;
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2540 = F_memcmp(m, v2537+v2527-v2529, int32(_a_F_italian_UTF_8_stem_27), v2529)
	mBase = m.M
	if v2540 != 0 {
		v2544 = v2531
		goto L624
	} else {
		goto L626
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2527 - v2529
	v2544 = int32(1)
	goto L624
L627:
	;
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2547
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2549)))
	if v2547 < v2550 {
		goto L61
	} else {
		goto L628
	}
L628:
	;
	v2552 = F_slice_del(m, l0)
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L8
	} else {
		goto L629
	}
L629:
	;
	if v2552 < int32(0) {
		v3055 = v2552
		goto L13
	} else {
		goto L630
	}
L630:
	;
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2556
	v2558 = int32(2)
	v2560 = int32(0)
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2556-v2563 < v2558 {
		v2573 = v2560
		goto L632
	} else {
		goto L633
	}
L631:
	;
	if v2573 == int32(0) {
		goto L61
	} else {
		goto L635
	}
L632:
	;
	goto L631
L633:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2569 = F_memcmp(m, v2566+v2556-v2558, int32(_a_F_italian_UTF_8_stem_28), v2558)
	mBase = m.M
	if v2569 != 0 {
		v2573 = v2560
		goto L632
	} else {
		goto L634
	}
L634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2556 - v2558
	v2573 = int32(1)
	goto L632
L635:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2576
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2579 = *(*int32)(unsafe.Add(mBase, uint32(v2578)))
	if v2576 < v2579 {
		goto L61
	} else {
		goto L636
	}
L636:
	;
	v2581 = F_slice_del(m, l0)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L8
	} else {
		goto L637
	}
L637:
	;
	if v2581 < int32(0) {
		v3055 = v2581
		goto L13
	} else {
		goto L638
	}
L638:
	;
	goto L61
L639:
	;
	if int32(0) <= v2587 {
		goto L62
	} else {
		goto L640
	}
L640:
	;
	v3055 = v2587
	goto L13
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2593
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2595
	v2602 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_29), int32(87))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L8
	} else {
		goto L642
	}
L642:
	;
	if v2602 != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2604
	v2606 = F_slice_del(m, l0)
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L8
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2598
	goto L61
L646:
	;
	if v2606 < int32(0) {
		v3055 = v2606
		goto L13
	} else {
		goto L647
	}
L647:
	;
	goto L645
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2785
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2785 <= v2789 {
		v2943 = v2787
		goto L682
	} else {
		goto L683
	}
L649:
	;
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2783
	v2785 = v2783
	v2787 = v2783
	goto L648
L650:
	;
	if v2747 != 0 {
		goto L649
	} else {
		goto L673
	}
L651:
	;
	v2747 = v2740
	goto L650
L652:
	;
	if v2616 <= v2632 {
		v2740 = int32(-1)
		goto L651
	} else {
		goto L654
	}
L653:
	;
	v2740 = int32(0)
	goto L651
L654:
	;
	v2649 = int32(1)
	v2650 = v2616 - v2649
	v2652 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2633+v2650))))
	v2654 = v2652 & int32(255)
	if base.B2i32(v2650 == v2632)|base.B2i32(int32(0) <= v2652) != 0 {
		v2712 = v2654
		v2716 = v2649
		goto L655
	} else {
		goto L656
	}
L655:
	;
	if int32(242) < v2712 {
		goto L663
	} else {
		goto L664
	}
L656:
	;
	v2661 = v2654 & int32(63)
	v2663 = v2616 - int32(2)
	v2665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2633+v2663))))
	v2667 = v2665 << (uint(int32(6)) % 32)
	if base.B2i32(v2663 != v2632)&base.B2i32(base.Ui32(v2665) < base.Ui32(int32(192))) == int32(0) {
		goto L657
	} else {
		goto L658
	}
L657:
	;
	v2712 = v2667&int32(1984) | v2661
	v2716 = int32(2)
	goto L655
L658:
	;
	goto L659
L659:
	;
	v2680 = v2667&int32(4032) | v2661
	v2682 = v2616 - int32(3)
	v2684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2633+v2682))))
	if base.B2i32(v2682 != v2632)&base.B2i32(base.Ui32(v2684) < base.Ui32(int32(224))) == int32(0) {
		goto L660
	} else {
		goto L661
	}
L660:
	;
	v2712 = v2684<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v2680
	v2716 = int32(3)
	goto L655
L661:
	;
	goto L662
L662:
	;
	v2702 = int32(4)
	v2704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2616+v2633-v2702))))
	v2712 = v2684<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_30) | v2704&int32(7)<<(uint(int32(18))%32) | v2680
	v2716 = v2702
	goto L655
L663:
	;
	v2747 = v2716
	goto L650
L664:
	;
	goto L665
L665:
	;
	v2718 = v2712 - int32(97)
	if v2718 < int32(0) {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v2747 = v2716
	goto L650
L667:
	;
	goto L668
L668:
	;
	v2724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2718)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[1]))))
	if int32(base.Ui32(v2724)>>(uint(v2718&int32(7))%32))&int32(1) == int32(0) {
		goto L669
	} else {
		goto L670
	}
L669:
	;
	v2747 = v2716
	goto L650
L670:
	;
	goto L671
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2616 - v2716
	goto L672
L672:
	;
	goto L653
L673:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2748
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2750)+8))
	if v2748 < v2751 {
		goto L649
	} else {
		goto L674
	}
L674:
	;
	v2753 = F_slice_del(m, l0)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L8
	} else {
		goto L675
	}
L675:
	;
	if v2753 < int32(0) {
		v3055 = v2753
		goto L13
	} else {
		goto L676
	}
L676:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2757
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2757 <= v2759 {
		goto L649
	} else {
		goto L677
	}
L677:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2761+v2757-int32(1)))))
	if v2765 != int32(105) {
		goto L649
	} else {
		goto L678
	}
L678:
	;
	v2769 = v2757 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2769
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2769
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2773 = *(*int32)(unsafe.Add(mBase, uint32(v2772)+8))
	if v2757 <= v2773 {
		goto L649
	} else {
		goto L679
	}
L679:
	;
	v2775 = F_slice_del(m, l0)
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		goto L8
	} else {
		goto L680
	}
L680:
	;
	if v2775 < int32(0) {
		v3055 = v2775
		goto L13
	} else {
		goto L681
	}
L681:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2785 = v2779
	v2787 = v2780
	goto L648
L682:
	;
	v2944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2944
	v2947 = v2944
	v2949 = v2943
	goto L713
L683:
	;
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2791+v2785-int32(1)))))
	if v2795 != int32(104) {
		v2943 = v2787
		goto L682
	} else {
		goto L684
	}
L684:
	;
	v2799 = v2785 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2799
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2816 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L688
L685:
	;
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2943 = v2940
	goto L682
L686:
	;
	if v2930 != 0 {
		goto L685
	} else {
		goto L709
	}
L687:
	;
	v2930 = v2923
	goto L686
L688:
	;
	if v2799 <= v2815 {
		v2923 = int32(-1)
		goto L687
	} else {
		goto L690
	}
L689:
	;
	v2923 = int32(0)
	goto L687
L690:
	;
	v2832 = int32(1)
	v2833 = v2799 - v2832
	v2835 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2816+v2833))))
	v2837 = v2835 & int32(255)
	if base.B2i32(v2833 == v2815)|base.B2i32(int32(0) <= v2835) != 0 {
		v2895 = v2837
		v2899 = v2832
		goto L691
	} else {
		goto L692
	}
L691:
	;
	if int32(103) < v2895 {
		goto L699
	} else {
		goto L700
	}
L692:
	;
	v2844 = v2837 & int32(63)
	v2846 = v2799 - int32(2)
	v2848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816+v2846))))
	v2850 = v2848 << (uint(int32(6)) % 32)
	if base.B2i32(v2846 != v2815)&base.B2i32(base.Ui32(v2848) < base.Ui32(int32(192))) == int32(0) {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v2895 = v2850&int32(1984) | v2844
	v2899 = int32(2)
	goto L691
L694:
	;
	goto L695
L695:
	;
	v2863 = v2850&int32(4032) | v2844
	v2865 = v2799 - int32(3)
	v2867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816+v2865))))
	if base.B2i32(v2865 != v2815)&base.B2i32(base.Ui32(v2867) < base.Ui32(int32(224))) == int32(0) {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	v2895 = v2867<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_10) | v2863
	v2899 = int32(3)
	goto L691
L697:
	;
	goto L698
L698:
	;
	v2885 = int32(4)
	v2887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2799+v2816-v2885))))
	v2895 = v2867<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_30) | v2887&int32(7)<<(uint(int32(18))%32) | v2863
	v2899 = v2885
	goto L691
L699:
	;
	v2930 = v2899
	goto L686
L700:
	;
	goto L701
L701:
	;
	v2901 = v2895 - int32(99)
	if v2901 < int32(0) {
		goto L702
	} else {
		goto L703
	}
L702:
	;
	v2930 = v2899
	goto L686
L703:
	;
	goto L704
L704:
	;
	v2907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2901)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[2]))))
	if int32(base.Ui32(v2907)>>(uint(v2901&int32(7))%32))&int32(1) == int32(0) {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v2930 = v2899
	goto L686
L706:
	;
	goto L707
L707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2799 - v2899
	goto L708
L708:
	;
	goto L689
L709:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(v2931)+8))
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2933 < v2932 {
		goto L685
	} else {
		goto L710
	}
L710:
	;
	v2935 = F_slice_del(m, l0)
	mBase = m.M
	v2936 = m.ExcPending
	if v2936 != 0 {
		goto L8
	} else {
		goto L711
	}
L711:
	;
	if v2935 < int32(0) {
		v3055 = v2935
		goto L13
	} else {
		goto L712
	}
L712:
	;
	goto L685
L713:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2947
	if v2949 <= v2947 {
		goto L718
	} else {
		goto L719
	}
L714:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2944
	v3055 = int32(1)
	goto L13
L715:
	;
	goto L714
L716:
	;
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2947 = v3051
	v2949 = v3050
	goto L713
L717:
	;
	v2991 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L731
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2947
	v2988 = v2947
	v2989 = v2949
	goto L717
L719:
	;
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2953+v2947))))
	v2957 = v2955 - int32(73)
	v2958 = int32(0)
	if base.B2i32(v2957 == v2958)|base.B2i32(v2957 == int32(12)) == v2958 {
		goto L718
	} else {
		goto L720
	}
L720:
	;
	v2967 = F_find_among(m, l0, int32(_a_F_italian_UTF_8_stem_31), int32(3))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L8
	} else {
		goto L721
	}
L721:
	;
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2969
	switch v2967 - int32(1) {
	case 0:
		goto L723
	case 1:
		goto L722
	case 2:
		goto L724
	default:
		goto L716
	}
L722:
	;
	v2982 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_32))
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L8
	} else {
		goto L727
	}
L723:
	;
	v2976 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_33))
	mBase = m.M
	v2977 = m.ExcPending
	if v2977 != 0 {
		goto L8
	} else {
		goto L725
	}
L724:
	;
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2988 = v2969
	v2989 = v2973
	goto L717
L725:
	;
	if int32(0) <= v2976 {
		goto L716
	} else {
		goto L726
	}
L726:
	;
	v3055 = v2976
	goto L13
L727:
	;
	if int32(0) <= v2982 {
		goto L716
	} else {
		goto L728
	}
L728:
	;
	v3055 = v2982
	goto L13
L729:
	;
	if v3043 < int32(0) {
		goto L715
	} else {
		goto L749
	}
L731:
	;
	goto L732
L732:
	;
	goto L733
L733:
	;
	v2998 = v2988
	v3000 = int32(1)
	goto L736
L735:
	;
	v3043 = v3028
	goto L729
L736:
	;
	if v2989 <= v2998 {
		goto L738
	} else {
		goto L739
	}
L737:
	;
	goto L735
L738:
	;
	v3043 = int32(-1)
	goto L729
L739:
	;
	goto L740
L740:
	;
	v3005 = v2998 + int32(1)
	v3007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2991+v2998))))
	if base.Ui32(v3007) < base.Ui32(int32(192)) {
		v3028 = v3005
		goto L741
	} else {
		goto L742
	}
L741:
	;
	v3029 = int32(1)
	if v3029 < v3000 {
		v2998 = v3028
		v3000 = v3000 - v3029
		goto L736
	} else {
		goto L748
	}
L742:
	;
	if v2989 <= v3005 {
		v3028 = v3005
		goto L741
	} else {
		goto L743
	}
L743:
	;
	v3014 = v3005
	goto L744
L744:
	;
	v3017 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2991+v3014))))
	if int32(-65) < v3017 {
		v3028 = v3014
		goto L741
	} else {
		goto L746
	}
L745:
	;
	v3028 = v2989
	goto L741
L746:
	;
	v3021 = v3014 + int32(1)
	if v3021 != v2989 {
		v3014 = v3021
		goto L744
	} else {
		goto L747
	}
L747:
	;
	goto L745
L748:
	;
	goto L737
L749:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3043
	goto L716
}
func F_iterate_values_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v4&int32(1) != 0 {
		v7 = F_pstrdup(m, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = F_strlen(m, v7)
			mBase = m.M
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v13].(func(*base.Module, int32, int32, int32))(m, v11, v7, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_ivfflatbuildempty(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = m.G0
	v5 = v3 - int32(176)
	m.G0 = v5
	v8 = F_BuildIndexInfo(m, l0)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_BuildIndex_2(m, int32(0), l0, v8, v5, int32(3))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v5 + int32(176)
			return
		}
	}
}
func F_ivfflatcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v99 float64
	_ = v99
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v21 != 0 {
		v22 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v22
		F_genericcostestimate(m, l0, l1, l2, v19+int32(16))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
			v45 = F_index_open(m, v43, int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				F_IvfflatGetMetaPageInfo(m, v45, v19+int32(12), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_relation_close(m, v45, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatcostestimate[0]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
						F_get_tablespace_page_costs(m, v59, int32(0), v19)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v63 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
							v66 = *(*float64)(unsafe.Add(mBase, uint32(v19)+64))
							v67 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
							v68 = base.F64_sub(v66, v67)
							v70 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
							v71 = base.F64_add(base.F64_mul(base.F64_mul(v63, float64(-0.5)), v68), v70)
							v72 = float64(1)
							v75 = base.F64_div(base.F64_convert_i32_s(v56), base.F64_convert_i32_s(v57))
							if base.F64_gt(v75, v72) != 0 {
								v78 = v72
							} else {
								v78 = v75
							}
							v79 = base.F64_mul(v71, v78)
							if base.F64_lt(v78, float64(0.5)) == int32(0) {
								v99 = v79
							} else {
								v84 = base.F64_mul(v63, v78)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+116))
								v88 = base.F64_convert_i32_u(v87)
								if base.F64_gt(v84, v88) == int32(0) {
									v99 = v79
								} else {
									v99 = base.F64_add(base.F64_mul(base.F64_sub(v88, v84), v67), base.F64_add(base.F64_mul(base.F64_mul(v84, float64(-0.5)), v68), v79))
								}
							}
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = v99
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v71
							v104 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
							*(*float64)(unsafe.Add(mBase, uint32(l5))) = v104
							v106 = *(*float64)(unsafe.Add(mBase, uint32(v19)+40))
							*(*float64)(unsafe.Add(mBase, uint32(l6))) = v106
							*(*float64)(unsafe.Add(mBase, uint32(l7))) = v63
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	} else {
		v109 = int64(9218868437227405312)
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v109
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v109
		v113 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l5))) = v113
		*(*int64)(unsafe.Add(mBase, uint32(l6))) = v113
		*(*int64)(unsafe.Add(mBase, uint32(l7))) = v113
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(2)
		m.G0 = v19 + int32(80)
		return
	}
}
