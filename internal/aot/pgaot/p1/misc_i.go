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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = v8 - int32(1)
		if base.Ui32(int32(3)) <= base.Ui32(v10) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
				F_errmsg_internal(m, int32(507893), v6)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(514767), int32(824), int32(20967))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_consts[747])))
			v20 = v17
			m.G0 = v6 + int32(16)
			return v20
		}
	} else {
		v20 = int32(3)
		m.G0 = v6 + int32(16)
		return v20
	}
}
func F_IdleSessionTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1204])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[48])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[497]))
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
	if base.Ui32(int32(16384)) <= base.Ui32(v15) {
		v23 = int32(1024)
	} else {
		v23 = int32(base.Ui32(v18) >> (uint(int32(4)) % 32))
	}
	*(*int32)(unsafe.Add(mBase, _consts[201])) = v23
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
		*(*int32)(unsafe.Add(mBase, _consts[779])) = v5
		v11 = *(*int32)(unsafe.Add(mBase, _consts[517]))
		F_AddWaitEventToSet(m, v5, int32(1), int32(-1), v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
			if v15 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[779]))
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return base.B2i32((v3-int32(7))&int32(-9) == int32(0))
}
func F__intbig_out(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(21513)
			F_errmsg(m, int32(202712), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(514868), int32(46), int32(73057))
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
	if base.Ui32(v2-int32(32768)) <= base.Ui32(int32(-65537)) {
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
				F_errmsg(m, int32(420378), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(384), int32(586682))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v2, v3, v4, v5, v6, v7, int32(1), v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
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
			F_errmsg(m, int32(450994), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(522708), int32(1599), int32(353999))
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 != int32(1) {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v50 = v48 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v50
		if v5 == int32(0) {
		} else {
			v54 = v50
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32)))) = int32(0)
		}
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v10+int32(1) < v9 {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v50 = v48 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v50
			if v5 == int32(0) {
			} else {
				v54 = v50
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32)))) = int32(0)
			}
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = v9 - int32(-64)
			v19 = F_repalloc(m, v14, v16*int32(10))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v19
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				v27 = F_repalloc(m, v24, v16<<(uint(int32(2))%32))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v27
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					v33 = F_repalloc(m, v32, v16)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v33
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v16
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v40 = int32(1)
						v41 = v39 + v40
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v41
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v43&v40 != 0 {
							v54 = v41
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32)))) = int32(0)
						} else {
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
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
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
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
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int64
	_ = v438
	var v440 int32
	_ = v440
	var v441 int64
	_ = v441
	var v445 int64
	_ = v445
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v452 int64
	_ = v452
	var v455 int64
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int64
	_ = v495
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int64
	_ = v585
	var v586 int32
	_ = v586
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v636 int64
	_ = v636
	var v639 int32
	_ = v639
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int64
	_ = v699
	var v704 int64
	_ = v704
	var v709 int32
	_ = v709
	var v712 int64
	_ = v712
	var v713 int32
	_ = v713
	var v723 int64
	_ = v723
	var v726 int64
	_ = v726
	var v728 int64
	_ = v728
	var v731 int64
	_ = v731
	var v732 int64
	_ = v732
	var v738 int64
	_ = v738
	var v739 int32
	_ = v739
	var v750 int64
	_ = v750
	var v752 int64
	_ = v752
	var v753 int32
	_ = v753
	var v755 int64
	_ = v755
	var v758 int64
	_ = v758
	var v759 int64
	_ = v759
	var v764 int64
	_ = v764
	var v765 int64
	_ = v765
	var v767 int64
	_ = v767
	var v769 int64
	_ = v769
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v773 int32
	_ = v773
	var v774 int64
	_ = v774
	var v775 int32
	_ = v775
	var v782 int64
	_ = v782
	var v783 int32
	_ = v783
	var v786 int64
	_ = v786
	var v787 int64
	_ = v787
	var v789 int64
	_ = v789
	var v791 int64
	_ = v791
	var v792 int32
	_ = v792
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int64
	_ = v813
	var v814 int64
	_ = v814
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int64
	_ = v834
	var v835 int64
	_ = v835
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int64
	_ = v855
	var v856 int64
	_ = v856
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int64
	_ = v874
	var v875 int64
	_ = v875
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int64
	_ = v895
	var v896 int64
	_ = v896
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int64
	_ = v916
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int64
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int64
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	v11 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(144)
	m.G0 = v25
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v11
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v696 - int32(21) {
	case 0:
		goto L202
	default:
		goto L201
	case 2:
		goto L203
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	v686 = v662
	v688 = v664
	v689 = v665
	v691 = v667
	v693 = v669
	goto L1
L3:
	;
	if v622 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L4:
	;
	if l3 != 0 {
		v622 = v598
		v624 = v600
		v626 = v602
		v627 = v603
		v629 = v605
		v631 = v607
		goto L3
	} else {
		goto L190
	}
L5:
	;
	if v578 == int32(0) {
		v598 = v568
		v600 = v570
		v602 = v572
		v603 = v573
		v605 = v575
		v607 = v577
		goto L4
	} else {
		goto L187
	}
L6:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+48)) = uint8(v554)
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	v568 = v551
	v570 = v552
	v572 = v406
	v573 = v407
	v575 = v409
	v577 = v411
	v578 = v412
	goto L5
L7:
	;
	v547 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+48)) = uint8(v547)
	v568 = v535
	v570 = v537
	v572 = v539
	v573 = v540
	v575 = v542
	v577 = v544
	v578 = v545
	goto L5
L8:
	;
	if l3 != 0 {
		goto L147
	} else {
		goto L148
	}
L9:
	;
	v54 = v11
	v56 = v11
	v57 = v11
	v58 = v11
	v59 = v11
	v60 = v11
	v61 = v11
	v62 = v11
	v63 = v11
	goto L17
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v31 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if l3 == int32(0) {
		v598 = v11
		v600 = v11
		v602 = v11
		v603 = v11
		v605 = v11
		v607 = v11
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v406 = v11
	v407 = v11
	v408 = v11
	v409 = v11
	v410 = v11
	v411 = v11
	v412 = v11
	v413 = v11
	goto L8
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(20)
	v535 = v11
	v537 = v11
	v539 = v11
	v540 = v11
	v542 = v11
	v544 = v11
	v545 = v11
	goto L7
L15:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L120
	} else {
		goto L136
	}
L16:
	;
	F_errorConflictingDefElem(m, v68, l0)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L120
	} else {
		goto L135
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v54<<(uint(int32(2))%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v70 != int32(97) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v331 = int32(398034)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, _consts[390])))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v335 == int32(0) {
		v354 = v334
		v355 = v335
		goto L124
	} else {
		goto L125
	}
L19:
	;
	goto L18
L20:
	;
	v328 = v54 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v328 < v329 {
		v54 = v328
		v56 = v319
		v57 = v320
		v58 = v321
		v59 = v322
		v60 = v323
		v61 = v324
		v62 = v325
		v63 = v326
		goto L17
	} else {
		goto L122
	}
L21:
	;
	v79 = int32(102341)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, _consts[391])))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v83 == int32(0) {
		v102 = v82
		v103 = v83
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v73 != int32(115) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+2)))
	if v76 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	if v63 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v77)
	v319 = v56
	v320 = v57
	v321 = v58
	v322 = v59
	v323 = v60
	v324 = v61
	v325 = v62
	v326 = v68
	goto L20
L26:
	;
	if v103-v102 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	goto L26
L28:
	;
	if v82 != v83 {
		v102 = v82
		v103 = v83
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v87 = v69
	v88 = v79
	goto L30
L30:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v92 == int32(0) {
		v102 = v91
		v103 = v92
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v102 = v91
	v103 = v92
	goto L27
L32:
	;
	v95 = int32(1)
	if v91 == v92 {
		v87 = v87 + v95
		v88 = v88 + v95
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v58 != 0 {
		goto L16
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v109 = int32(89102)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _consts[392])))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v113 == int32(0) {
		v132 = v112
		v133 = v113
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v107)
	v319 = v56
	v320 = v57
	v321 = v68
	v322 = v59
	v323 = v60
	v324 = v61
	v325 = v62
	v326 = v63
	goto L20
L38:
	;
	if v133-v132 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	goto L38
L40:
	;
	if v112 != v113 {
		v132 = v112
		v133 = v113
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v117 = v69
	v118 = v109
	goto L42
L42:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v122 == int32(0) {
		v132 = v121
		v133 = v122
		goto L39
	} else {
		goto L44
	}
L43:
	;
	v132 = v121
	v133 = v122
	goto L39
L44:
	;
	v125 = int32(1)
	if v121 == v122 {
		v117 = v117 + v125
		v118 = v118 + v125
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v59 != 0 {
		goto L16
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v139 = int32(88319)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, _consts[393])))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v143 == int32(0) {
		v162 = v142
		v163 = v143
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v137 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v137)
	v319 = v56
	v320 = v57
	v321 = v58
	v322 = v68
	v323 = v60
	v324 = v61
	v325 = v62
	v326 = v63
	goto L20
L50:
	;
	if v163-v162 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	goto L50
L52:
	;
	if v142 != v143 {
		v162 = v142
		v163 = v143
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v147 = v69
	v148 = v139
	goto L54
L54:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v152 == int32(0) {
		v162 = v151
		v163 = v152
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v162 = v151
	v163 = v152
	goto L51
L56:
	;
	v155 = int32(1)
	if v151 == v152 {
		v147 = v147 + v155
		v148 = v148 + v155
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v56 != 0 {
		goto L16
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v169 = int32(361562)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, _consts[394])))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v173 == int32(0) {
		v192 = v172
		v193 = v173
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v167)
	v319 = v68
	v320 = v57
	v321 = v58
	v322 = v59
	v323 = v60
	v324 = v61
	v325 = v62
	v326 = v63
	goto L20
L62:
	;
	if v193-v192 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	goto L62
L64:
	;
	if v172 != v173 {
		v192 = v172
		v193 = v173
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v177 = v69
	v178 = v169
	goto L66
L66:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v182 == int32(0) {
		v192 = v181
		v193 = v182
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v192 = v181
	v193 = v182
	goto L63
L68:
	;
	v185 = int32(1)
	if v181 == v182 {
		v177 = v177 + v185
		v178 = v178 + v185
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	if v62 != 0 {
		goto L16
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v199 = int32(361571)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _consts[395])))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v203 == int32(0) {
		v222 = v202
		v223 = v203
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v197)
	v319 = v56
	v320 = v57
	v321 = v58
	v322 = v59
	v323 = v60
	v324 = v61
	v325 = v68
	v326 = v63
	goto L20
L74:
	;
	if v223-v222 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	goto L74
L76:
	;
	if v202 != v203 {
		v222 = v202
		v223 = v203
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v207 = v69
	v208 = v199
	goto L78
L78:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	if v212 == int32(0) {
		v222 = v211
		v223 = v212
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v222 = v211
	v223 = v212
	goto L75
L80:
	;
	v215 = int32(1)
	if v211 == v212 {
		v207 = v207 + v215
		v208 = v208 + v215
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if v57 != 0 {
		goto L16
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v229 = int32(418134)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _consts[396])))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v233 == int32(0) {
		v252 = v232
		v253 = v233
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v227)
	v319 = v56
	v320 = v68
	v321 = v58
	v322 = v59
	v323 = v60
	v324 = v61
	v325 = v62
	v326 = v63
	goto L20
L86:
	;
	if v253-v252 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	goto L86
L88:
	;
	if v232 != v233 {
		v252 = v232
		v253 = v233
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v237 = v69
	v238 = v229
	goto L90
L90:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	if v242 == int32(0) {
		v252 = v241
		v253 = v242
		goto L87
	} else {
		goto L92
	}
L91:
	;
	v252 = v241
	v253 = v242
	goto L87
L92:
	;
	v245 = int32(1)
	if v241 == v242 {
		v237 = v237 + v245
		v238 = v238 + v245
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	if v61 != 0 {
		goto L16
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v259 = int32(408612)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, _consts[397])))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v263 == int32(0) {
		v282 = v262
		v283 = v263
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v257 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v257)
	v319 = v56
	v320 = v57
	v321 = v58
	v322 = v59
	v323 = v60
	v324 = v68
	v325 = v62
	v326 = v63
	goto L20
L98:
	;
	if v283-v282 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L99:
	;
	goto L98
L100:
	;
	if v262 != v263 {
		v282 = v262
		v283 = v263
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v267 = v69
	v268 = v259
	goto L102
L102:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	if v272 == int32(0) {
		v282 = v271
		v283 = v272
		goto L99
	} else {
		goto L104
	}
L103:
	;
	v282 = v271
	v283 = v272
	goto L99
L104:
	;
	v275 = int32(1)
	if v271 == v272 {
		v267 = v267 + v275
		v268 = v268 + v275
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	if v60 != 0 {
		goto L16
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v289 = int32(24407)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, _consts[398])))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v293 == int32(0) {
		v312 = v292
		v313 = v293
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v287 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v287)
	v319 = v56
	v320 = v57
	v321 = v58
	v322 = v59
	v323 = v68
	v324 = v61
	v325 = v62
	v326 = v63
	goto L20
L110:
	;
	if v313-v312 != 0 {
		goto L19
	} else {
		goto L118
	}
L111:
	;
	goto L110
L112:
	;
	if v292 != v293 {
		v312 = v292
		v313 = v293
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v297 = v69
	v298 = v289
	goto L114
L114:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v302 == int32(0) {
		v312 = v301
		v313 = v302
		goto L111
	} else {
		goto L116
	}
L115:
	;
	v312 = v301
	v313 = v302
	goto L111
L116:
	;
	v305 = int32(1)
	if v301 == v302 {
		v297 = v297 + v305
		v298 = v298 + v305
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	if v315 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	v316 = F_defGetQualifiedName(m, v68)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	return
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v316
	v319 = v56
	v320 = v57
	v321 = v58
	v322 = v59
	v323 = v60
	v324 = v61
	v325 = v62
	v326 = v63
	goto L20
L122:
	;
	v406 = v319
	v407 = v320
	v408 = v321
	v409 = v322
	v410 = v323
	v411 = v324
	v412 = v325
	v413 = v326
	goto L8
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L120
	} else {
		goto L131
	}
L124:
	;
	goto L123
L125:
	;
	if v334 != v335 {
		v354 = v334
		v355 = v335
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v339 = v69
	v340 = v331
	goto L127
L127:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)))
	if v344 == int32(0) {
		v354 = v343
		v355 = v344
		goto L124
	} else {
		goto L129
	}
L128:
	;
	v354 = v343
	v355 = v344
	goto L124
L129:
	;
	v347 = int32(1)
	if v343 == v344 {
		v339 = v339 + v347
		v340 = v340 + v347
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	if v355-v354 == int32(0) {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+128)) = v363
	F_errmsg_internal(m, int32(458203), v25+int32(128))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L120
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(523122), int32(1362), int32(160295))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L120
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg(m, int32(564838), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L120
	} else {
		goto L137
	}
L137:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	F_parser_errposition(m, l0, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L120
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(523122), int32(1358), int32(160295))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L120
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(1)
	v523 = int32(0)
	if v410 != 0 {
		v551 = v523
		v552 = v523
		goto L6
	} else {
		goto L186
	}
L141:
	;
	v495 = F_defGetInt64(m, v408)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L120
	} else {
		goto L176
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L120
	} else {
		goto L169
	}
L143:
	;
	if v410 == int32(0) {
		v568 = v470
		v570 = v471
		v572 = v406
		v573 = v407
		v575 = v409
		v577 = v411
		v578 = v412
		goto L5
	} else {
		goto L168
	}
L144:
	;
	v467 = int32(0)
	if v408 != 0 {
		v493 = v467
		v494 = v467
		goto L141
	} else {
		goto L167
	}
L145:
	;
	v465 = int32(0)
	v493 = v465
	v494 = v465
	goto L141
L146:
	;
	v420 = F_defGetTypeName(m, v413)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L120
	} else {
		goto L153
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	if v413 != 0 {
		goto L146
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	if v413 == int32(0) {
		goto L144
	} else {
		goto L152
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(20)
	if v408 != 0 {
		goto L145
	} else {
		goto L151
	}
L151:
	;
	goto L140
L152:
	;
	goto L146
L153:
	;
	v422 = F_typenameTypeId(m, l0, v420)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L120
	} else {
		goto L154
	}
L154:
	;
	if base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v422-int32(20)))&base.B2i32(v422 != int32(23)) != 0 {
		goto L142
	} else {
		goto L155
	}
L155:
	;
	if l3 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v433 = int32(0)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v435 - int32(20) {
	case 0:
		goto L160
	case 1:
		goto L162
	default:
		v458 = v433
		v459 = v433
		goto L159
	case 3:
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v422
	if v408 == int32(0) {
		goto L140
	} else {
		goto L166
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v422
	if v408 != 0 {
		v493 = v458
		v494 = v459
		goto L141
	} else {
		goto L165
	}
L160:
	;
	v452 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v455 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v458 = base.B2i32(v455 == int64(9223372036854775807))
	v459 = base.B2i32(v452 == int64(-9223372036854775807-1))
	goto L159
L161:
	;
	v445 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v447 = base.B2i32(v445 == int64(2147483647))
	v448 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v448 != int64(-2147483648) {
		v458 = v447
		v459 = v433
		goto L159
	} else {
		goto L164
	}
L162:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v440 = base.B2i32(v438 == int64(32767))
	v441 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v441 != int64(-32768) {
		v458 = v440
		v459 = v433
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v458 = v440
	v459 = int32(1)
	goto L159
L164:
	;
	v458 = v447
	v459 = int32(1)
	goto L159
L165:
	;
	v470 = v458
	v471 = v459
	goto L143
L166:
	;
	goto L145
L167:
	;
	v470 = v467
	v471 = v467
	goto L143
L168:
	;
	v551 = v470
	v552 = v471
	goto L6
L169:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L120
	} else {
		goto L170
	}
L170:
	;
	if l2 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v483 = int32(95524)
	goto L173
L172:
	;
	v483 = int32(95582)
	goto L173
L173:
	;
	F_errmsg(m, v483, int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L120
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(523122), int32(1384), int32(160295))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L120
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
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = v495
	if v495 != int64(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	if v410 != 0 {
		v551 = v493
		v552 = v494
		goto L6
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L120
	} else {
		goto L182
	}
L180:
	;
	if l3 == int32(0) {
		v568 = v493
		v570 = v494
		v572 = v406
		v573 = v407
		v575 = v409
		v577 = v411
		v578 = v412
		goto L5
	} else {
		goto L181
	}
L181:
	;
	v535 = v493
	v537 = v494
	v539 = v406
	v540 = v407
	v542 = v409
	v544 = v411
	v545 = v412
	goto L7
L182:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L120
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(251916), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L120
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(523122), int32(1418), int32(160295))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L120
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
	v535 = v523
	v537 = v523
	v539 = v406
	v540 = v407
	v542 = v409
	v544 = v411
	v545 = v412
	goto L7
L187:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	if v582 == int32(0) {
		v622 = v568
		v624 = v570
		v626 = v572
		v627 = v573
		v629 = v575
		v631 = v577
		goto L3
	} else {
		goto L188
	}
L188:
	;
	v585 = F_defGetInt64(m, v578)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L120
	} else {
		goto L189
	}
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = v585
	v662 = v570
	v664 = v572
	v665 = v573
	v667 = v575
	v669 = v577
	goto L2
L190:
	;
	if v598 == int32(0) {
		v686 = v600
		v688 = v602
		v689 = v603
		v691 = v605
		v693 = v607
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v622 = v598
	v624 = v600
	v626 = v602
	v627 = v603
	v629 = v605
	v631 = v607
	goto L3
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(-1)
	v662 = v624
	v664 = v626
	v665 = v627
	v667 = v629
	v669 = v631
	goto L2
L193:
	;
	v636 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if v636 <= int64(0) {
		goto L192
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v639 - int32(21) {
	case 0:
		goto L199
	default:
		goto L197
	case 2:
		goto L198
	}
L196:
	;
	goto L195
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(9223372036854775807)
	v662 = v624
	v664 = v626
	v665 = v627
	v667 = v629
	v669 = v631
	goto L2
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(2147483647)
	v662 = v624
	v664 = v626
	v665 = v627
	v667 = v629
	v669 = v631
	goto L2
L199:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(32767)
	v662 = v624
	v664 = v626
	v665 = v627
	v667 = v629
	v669 = v631
	goto L2
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L120
	} else {
		goto L304
	}
L201:
	;
	if v689 != 0 {
		goto L209
	} else {
		goto L210
	}
L202:
	;
	v704 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if base.Ui64(v704-int64(32768)) < base.Ui64(int64(-65536)) {
		goto L200
	} else {
		goto L205
	}
L203:
	;
	v699 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if base.Ui64(int64(-4294967297)) < base.Ui64(v699-int64(2147483648)) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	goto L201
L206:
	;
	v738 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v739 - int32(21) {
	case 0:
		goto L229
	default:
		goto L228
	case 2:
		goto L230
	}
L207:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	goto L206
L208:
	;
	if v696 == int32(23) {
		goto L215
	} else {
		goto L216
	}
L209:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v689)+12))
	if v709 == int32(0) {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	if l3|v686 != int32(1) {
		goto L206
	} else {
		goto L214
	}
L212:
	;
	v712 = F_defGetInt64(m, v689)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L120
	} else {
		goto L213
	}
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = v712
	goto L207
L214:
	;
	goto L208
L215:
	;
	v723 = int64(-2147483648)
	goto L217
L216:
	;
	v723 = int64(-9223372036854775807 - 1)
	goto L217
L217:
	;
	if v696 == int32(21) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v726 = int64(-32768)
	goto L220
L219:
	;
	v726 = v723
	goto L220
L220:
	;
	v728 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if int64(0) <= v728 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v731 = int64(1)
	goto L223
L222:
	;
	v731 = v726
	goto L223
L223:
	;
	if v686 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v732 = v726
	goto L226
L225:
	;
	v732 = v731
	goto L226
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = v732
	goto L207
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L120
	} else {
		goto L299
	}
L228:
	;
	v750 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v738 < v750 {
		goto L238
	} else {
		goto L239
	}
L229:
	;
	if base.Ui64(v738-int64(32768)) < base.Ui64(int64(-65536)) {
		goto L227
	} else {
		goto L232
	}
L230:
	;
	if base.Ui64(int64(-4294967297)) < base.Ui64(v738-int64(2147483648)) {
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L227
L232:
	;
	goto L228
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L120
	} else {
		goto L295
	}
L234:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L120
	} else {
		goto L291
	}
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L120
	} else {
		goto L287
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L120
	} else {
		goto L283
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L120
	} else {
		goto L279
	}
L238:
	;
	if v691 != 0 {
		goto L243
	} else {
		goto L244
	}
L239:
	;
	goto L240
L240:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L120
	} else {
		goto L275
	}
L241:
	;
	v769 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v769 < v767 {
		goto L236
	} else {
		goto L254
	}
L242:
	;
	if v764 < v765 {
		goto L237
	} else {
		goto L253
	}
L243:
	;
	v752 = F_defGetInt64(m, v691)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L120
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if l3 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v752
	v755 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v764 = v752
	v765 = v755
	goto L242
L247:
	;
	v758 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v764 = v758
	v765 = v738
	goto L242
L248:
	;
	goto L249
L249:
	;
	v759 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if int64(0) < v759 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v738
	v767 = v738
	goto L241
L251:
	;
	goto L252
L252:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v750
	v764 = v750
	v765 = v738
	goto L242
L253:
	;
	v767 = v764
	goto L241
L254:
	;
	if v688 != 0 {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v787 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v786 < v787 {
		goto L235
	} else {
		goto L266
	}
L256:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v688)+12))
	if v771 != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	if l3 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L259:
	;
	v772 = F_defGetInt64(m, v688)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L120
	} else {
		goto L262
	}
L260:
	;
	v774 = v767
	goto L261
L261:
	;
	v775 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)) = uint8(v775)
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v774
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	v786 = v774
	goto L255
L262:
	;
	v774 = v772
	goto L261
L263:
	;
	v782 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v786 = v782
	goto L255
L264:
	;
	goto L265
L265:
	;
	v783 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)) = uint8(v783)
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v767
	v786 = v767
	goto L255
L266:
	;
	v789 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v789 < v786 {
		goto L234
	} else {
		goto L267
	}
L267:
	;
	if v693 != 0 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	m.G0 = v25 + int32(144)
	return
L269:
	;
	v791 = F_defGetInt64(m, v693)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L120
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	if l3 == int32(0) {
		goto L268
	} else {
		goto L274
	}
L272:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+40)) = v791
	if v791 <= int64(0) {
		goto L233
	} else {
		goto L273
	}
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	goto L268
L274:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+40)) = int64(1)
	goto L268
L275:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L120
	} else {
		goto L276
	}
L276:
	;
	v813 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v814 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+88)) = v814
	*(*int64)(unsafe.Add(mBase, uint32(v25)+80)) = v813
	F_errmsg(m, int32(707093), v25+int32(80))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L120
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(523122), int32(1508), int32(160295))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L120
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L120
	} else {
		goto L280
	}
L280:
	;
	v834 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v835 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+72)) = v835
	*(*int64)(unsafe.Add(mBase, uint32(v25)+64)) = v834
	F_errmsg(m, int32(707205), v25-int32(-64))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L120
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(523122), int32(1529), int32(160295))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L120
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L120
	} else {
		goto L284
	}
L284:
	;
	v855 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v856 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v856
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v855
	F_errmsg(m, int32(707145), v25)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L120
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(523122), int32(1535), int32(160295))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L120
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L120
	} else {
		goto L288
	}
L288:
	;
	v874 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v875 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v875
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v874
	F_errmsg(m, int32(707203), v25+int32(16))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L120
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(523122), int32(1559), int32(160295))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L120
	} else {
		goto L290
	}
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L120
	} else {
		goto L292
	}
L292:
	;
	v895 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v896 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v896
	*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v895
	F_errmsg(m, int32(707143), v25+int32(32))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L120
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(523122), int32(1565), int32(160295))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L120
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
	v915 = m.ExcPending
	if v915 != 0 {
		goto L120
	} else {
		goto L296
	}
L296:
	;
	v916 = *(*int64)(unsafe.Add(mBase, uint32(l4)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v916
	F_errmsg(m, int32(251725), v25+int32(48))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L120
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(523122), int32(1575), int32(160295))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L120
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L120
	} else {
		goto L300
	}
L300:
	;
	v935 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v937 = F_format_type_be(m, v936)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L120
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = v937
	*(*int64)(unsafe.Add(mBase, uint32(v25)+96)) = v935
	F_errmsg(m, int32(203959), v25+int32(96))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L120
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(523122), int32(1500), int32(160295))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L120
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L120
	} else {
		goto L305
	}
L305:
	;
	v958 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v960 = F_format_type_be(m, v959)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L120
	} else {
		goto L306
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+120)) = v960
	*(*int64)(unsafe.Add(mBase, uint32(v25)+112)) = v958
	F_errmsg(m, int32(203901), v25+int32(112))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L120
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(523122), int32(1468), int32(160295))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L120
	} else {
		goto L308
	}
L308:
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
		v5 = *(*int32)(unsafe.Add(mBase, _consts[408]))
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	v6 = l5
	if l1 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1468))
		if v14 != v16 {
			v21 = F_LWLockAcquire(m, v15+int32(1476), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v25+int32(1476))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v33 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
						v38 = l0 + v33*int32(20) + int32(12)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						if v39 != 0 {
							v43 = v39
							v50 = v43 + l1&int32(134217727)
							v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_consts[1232]))))
							v57 = l2 + int32(20)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							if v58 != 0 {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1468))
								if v59 != v61 {
									v66 = F_LWLockAcquire(m, v60+int32(1476), int32(0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v70+int32(1476))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
												v83 = l0 + v78*int32(20) + int32(12)
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
												if v84 != 0 {
													v88 = v84
													*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
													v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													v99 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
													*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
													*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
													*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
													switch v6 {
													case 0:
														v109 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
														v112 = base.I32_div_u_s(int32(4096), v55)
														v117 = v112 - v109
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
														v119 = v117
													case 1:
														v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
														v119 = v108
													default:
														v116 = base.I32_div_u_s(int32(65536), v55)
														v117 = v116
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
														v119 = v117
													}
													v120 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
													v123 = int32(65535)
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
													return
												} else {
													v85 = F_get_segment_by_index(m, l0, v78)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v88 = v87
														*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
														v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
														v99 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
														*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
														*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
														*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
														switch v6 {
														case 0:
															v109 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
															v112 = base.I32_div_u_s(int32(4096), v55)
															v117 = v112 - v109
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
															v119 = v117
														case 1:
															v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
															v119 = v108
														default:
															v116 = base.I32_div_u_s(int32(65536), v55)
															v117 = v116
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
															v119 = v117
														}
														v120 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
														v123 = int32(65535)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
														return
													}
												}
											}
										}
									}
								} else {
									v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
									v83 = l0 + v78*int32(20) + int32(12)
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
									if v84 != 0 {
										v88 = v84
										*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										v99 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
										*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
										*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
										*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
										switch v6 {
										case 0:
											v109 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
											v112 = base.I32_div_u_s(int32(4096), v55)
											v117 = v112 - v109
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
											v119 = v117
										case 1:
											v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
											v119 = v108
										default:
											v116 = base.I32_div_u_s(int32(65536), v55)
											v117 = v116
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
											v119 = v117
										}
										v120 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
										v123 = int32(65535)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
										return
									} else {
										v85 = F_get_segment_by_index(m, l0, v78)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
											v88 = v87
											*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											v99 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
											*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
											switch v6 {
											case 0:
												v109 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
												v112 = base.I32_div_u_s(int32(4096), v55)
												v117 = v112 - v109
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											case 1:
												v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
												v119 = v108
											default:
												v116 = base.I32_div_u_s(int32(65536), v55)
												v117 = v116
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											}
											v120 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
											v123 = int32(65535)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
											return
										}
									}
								}
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v99 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
								switch v6 {
								case 0:
									v109 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
									v112 = base.I32_div_u_s(int32(4096), v55)
									v117 = v112 - v109
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								case 1:
									v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
									v119 = v108
								default:
									v116 = base.I32_div_u_s(int32(65536), v55)
									v117 = v116
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								}
								v120 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
								v123 = int32(65535)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
								return
							}
						} else {
							v40 = F_get_segment_by_index(m, l0, v33)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
								v43 = v42
								v50 = v43 + l1&int32(134217727)
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_consts[1232]))))
								v57 = l2 + int32(20)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								if v58 != 0 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1468))
									if v59 != v61 {
										v66 = F_LWLockAcquire(m, v60+int32(1476), int32(0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v70+int32(1476))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
													v83 = l0 + v78*int32(20) + int32(12)
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v84 != 0 {
														v88 = v84
														*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
														v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
														v99 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
														*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
														*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
														*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
														switch v6 {
														case 0:
															v109 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
															v112 = base.I32_div_u_s(int32(4096), v55)
															v117 = v112 - v109
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
															v119 = v117
														case 1:
															v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
															v119 = v108
														default:
															v116 = base.I32_div_u_s(int32(65536), v55)
															v117 = v116
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
															v119 = v117
														}
														v120 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
														v123 = int32(65535)
														*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
														return
													} else {
														v85 = F_get_segment_by_index(m, l0, v78)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v88 = v87
															*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
															v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
															v99 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
															*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
															*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
															*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
															*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
															switch v6 {
															case 0:
																v109 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
																v112 = base.I32_div_u_s(int32(4096), v55)
																v117 = v112 - v109
																*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
																v119 = v117
															case 1:
																v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
																v119 = v108
															default:
																v116 = base.I32_div_u_s(int32(65536), v55)
																v117 = v116
																*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
																v119 = v117
															}
															v120 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
															v123 = int32(65535)
															*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
															return
														}
													}
												}
											}
										}
									} else {
										v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
										v83 = l0 + v78*int32(20) + int32(12)
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
										if v84 != 0 {
											v88 = v84
											*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											v99 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
											*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
											switch v6 {
											case 0:
												v109 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
												v112 = base.I32_div_u_s(int32(4096), v55)
												v117 = v112 - v109
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											case 1:
												v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
												v119 = v108
											default:
												v116 = base.I32_div_u_s(int32(65536), v55)
												v117 = v116
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											}
											v120 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
											v123 = int32(65535)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
											return
										} else {
											v85 = F_get_segment_by_index(m, l0, v78)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
												v88 = v87
												*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
												v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												v99 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
												*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
												*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
												*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
												switch v6 {
												case 0:
													v109 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
													v112 = base.I32_div_u_s(int32(4096), v55)
													v117 = v112 - v109
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
													v119 = v117
												case 1:
													v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
													v119 = v108
												default:
													v116 = base.I32_div_u_s(int32(65536), v55)
													v117 = v116
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
													v119 = v117
												}
												v120 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
												v123 = int32(65535)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
												return
											}
										}
									}
								} else {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
									v99 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
									switch v6 {
									case 0:
										v109 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
										v112 = base.I32_div_u_s(int32(4096), v55)
										v117 = v112 - v109
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
										v119 = v117
									case 1:
										v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
										v119 = v108
									default:
										v116 = base.I32_div_u_s(int32(65536), v55)
										v117 = v116
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
										v119 = v117
									}
									v120 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
									v123 = int32(65535)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v33 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
			v38 = l0 + v33*int32(20) + int32(12)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
			if v39 != 0 {
				v43 = v39
				v50 = v43 + l1&int32(134217727)
				v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_consts[1232]))))
				v57 = l2 + int32(20)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				if v58 != 0 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1468))
					if v59 != v61 {
						v66 = F_LWLockAcquire(m, v60+int32(1476), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v70+int32(1476))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
									v83 = l0 + v78*int32(20) + int32(12)
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
									if v84 != 0 {
										v88 = v84
										*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										v99 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
										*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
										*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
										*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
										switch v6 {
										case 0:
											v109 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
											v112 = base.I32_div_u_s(int32(4096), v55)
											v117 = v112 - v109
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
											v119 = v117
										case 1:
											v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
											v119 = v108
										default:
											v116 = base.I32_div_u_s(int32(65536), v55)
											v117 = v116
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
											v119 = v117
										}
										v120 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
										v123 = int32(65535)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
										return
									} else {
										v85 = F_get_segment_by_index(m, l0, v78)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
											v88 = v87
											*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											v99 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
											*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
											switch v6 {
											case 0:
												v109 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
												v112 = base.I32_div_u_s(int32(4096), v55)
												v117 = v112 - v109
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											case 1:
												v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
												v119 = v108
											default:
												v116 = base.I32_div_u_s(int32(65536), v55)
												v117 = v116
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											}
											v120 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
											v123 = int32(65535)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
											return
										}
									}
								}
							}
						}
					} else {
						v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
						v83 = l0 + v78*int32(20) + int32(12)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						if v84 != 0 {
							v88 = v84
							*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							v99 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
							*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
							switch v6 {
							case 0:
								v109 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
								v112 = base.I32_div_u_s(int32(4096), v55)
								v117 = v112 - v109
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
								v119 = v117
							case 1:
								v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
								v119 = v108
							default:
								v116 = base.I32_div_u_s(int32(65536), v55)
								v117 = v116
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
								v119 = v117
							}
							v120 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
							v123 = int32(65535)
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
							return
						} else {
							v85 = F_get_segment_by_index(m, l0, v78)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
								v88 = v87
								*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v99 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
								switch v6 {
								case 0:
									v109 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
									v112 = base.I32_div_u_s(int32(4096), v55)
									v117 = v112 - v109
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								case 1:
									v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
									v119 = v108
								default:
									v116 = base.I32_div_u_s(int32(65536), v55)
									v117 = v116
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								}
								v120 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
								v123 = int32(65535)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
								return
							}
						}
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v99 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
					switch v6 {
					case 0:
						v109 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
						v112 = base.I32_div_u_s(int32(4096), v55)
						v117 = v112 - v109
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
						v119 = v117
					case 1:
						v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
						v119 = v108
					default:
						v116 = base.I32_div_u_s(int32(65536), v55)
						v117 = v116
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
						v119 = v117
					}
					v120 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
					v123 = int32(65535)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
					return
				}
			} else {
				v40 = F_get_segment_by_index(m, l0, v33)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v43 = v42
					v50 = v43 + l1&int32(134217727)
					v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_consts[1232]))))
					v57 = l2 + int32(20)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v58 != 0 {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1468))
						if v59 != v61 {
							v66 = F_LWLockAcquire(m, v60+int32(1476), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v70+int32(1476))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
										v83 = l0 + v78*int32(20) + int32(12)
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
										if v84 != 0 {
											v88 = v84
											*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											v99 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
											*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
											switch v6 {
											case 0:
												v109 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
												v112 = base.I32_div_u_s(int32(4096), v55)
												v117 = v112 - v109
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											case 1:
												v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
												v119 = v108
											default:
												v116 = base.I32_div_u_s(int32(65536), v55)
												v117 = v116
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
												v119 = v117
											}
											v120 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
											v123 = int32(65535)
											*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
											return
										} else {
											v85 = F_get_segment_by_index(m, l0, v78)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
												v88 = v87
												*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
												v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												v99 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
												*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
												*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
												*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
												switch v6 {
												case 0:
													v109 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
													v112 = base.I32_div_u_s(int32(4096), v55)
													v117 = v112 - v109
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
													v119 = v117
												case 1:
													v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
													v119 = v108
												default:
													v116 = base.I32_div_u_s(int32(65536), v55)
													v117 = v116
													*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
													v119 = v117
												}
												v120 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
												v123 = int32(65535)
												*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
												return
											}
										}
									}
								}
							}
						} else {
							v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
							v83 = l0 + v78*int32(20) + int32(12)
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
							if v84 != 0 {
								v88 = v84
								*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v99 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
								switch v6 {
								case 0:
									v109 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
									v112 = base.I32_div_u_s(int32(4096), v55)
									v117 = v112 - v109
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								case 1:
									v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
									v119 = v108
								default:
									v116 = base.I32_div_u_s(int32(65536), v55)
									v117 = v116
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								}
								v120 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
								v123 = int32(65535)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
								return
							} else {
								v85 = F_get_segment_by_index(m, l0, v78)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
									v88 = v87
									*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
									v99 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
									switch v6 {
									case 0:
										v109 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
										v112 = base.I32_div_u_s(int32(4096), v55)
										v117 = v112 - v109
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
										v119 = v117
									case 1:
										v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
										v119 = v108
									default:
										v116 = base.I32_div_u_s(int32(65536), v55)
										v117 = v116
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
										v119 = v117
									}
									v120 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
									v123 = int32(65535)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
									return
								}
							}
						}
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v99 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
						switch v6 {
						case 0:
							v109 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
							v112 = base.I32_div_u_s(int32(4096), v55)
							v117 = v112 - v109
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
							v119 = v117
						case 1:
							v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
							v119 = v108
						default:
							v116 = base.I32_div_u_s(int32(65536), v55)
							v117 = v116
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
							v119 = v117
						}
						v120 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
						v123 = int32(65535)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
						return
					}
				}
			}
		}
	} else {
		v50 = int32(0)
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_consts[1232]))))
		v57 = l2 + int32(20)
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
		if v58 != 0 {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1468))
			if v59 != v61 {
				v66 = F_LWLockAcquire(m, v60+int32(1476), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v70+int32(1476))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
							v83 = l0 + v78*int32(20) + int32(12)
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
							if v84 != 0 {
								v88 = v84
								*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v99 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
								switch v6 {
								case 0:
									v109 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
									v112 = base.I32_div_u_s(int32(4096), v55)
									v117 = v112 - v109
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								case 1:
									v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
									v119 = v108
								default:
									v116 = base.I32_div_u_s(int32(65536), v55)
									v117 = v116
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
									v119 = v117
								}
								v120 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
								v123 = int32(65535)
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
								return
							} else {
								v85 = F_get_segment_by_index(m, l0, v78)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
									v88 = v87
									*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
									v99 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
									switch v6 {
									case 0:
										v109 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
										v112 = base.I32_div_u_s(int32(4096), v55)
										v117 = v112 - v109
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
										v119 = v117
									case 1:
										v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
										v119 = v108
									default:
										v116 = base.I32_div_u_s(int32(65536), v55)
										v117 = v116
										*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
										v119 = v117
									}
									v120 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
									v123 = int32(65535)
									*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
									return
								}
							}
						}
					}
				}
			} else {
				v78 = int32(base.Ui32(v58) >> (uint(int32(27)) % 32))
				v83 = l0 + v78*int32(20) + int32(12)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
				if v84 != 0 {
					v88 = v84
					*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v99 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
					switch v6 {
					case 0:
						v109 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
						v112 = base.I32_div_u_s(int32(4096), v55)
						v117 = v112 - v109
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
						v119 = v117
					case 1:
						v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
						v119 = v108
					default:
						v116 = base.I32_div_u_s(int32(65536), v55)
						v117 = v116
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
						v119 = v117
					}
					v120 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
					v123 = int32(65535)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
					return
				} else {
					v85 = F_get_segment_by_index(m, l0, v78)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						v88 = v87
						*(*int32)(unsafe.Add(mBase, uint32(v88+v58&int32(134217727))+4)) = l1
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v99 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
						switch v6 {
						case 0:
							v109 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
							v112 = base.I32_div_u_s(int32(4096), v55)
							v117 = v112 - v109
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
							v119 = v117
						case 1:
							v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
							v119 = v108
						default:
							v116 = base.I32_div_u_s(int32(65536), v55)
							v117 = v116
							*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
							v119 = v117
						}
						v120 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
						v123 = int32(65535)
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
						return
					}
				}
			}
		} else {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v50))) = l2 - v95
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			v99 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v99
			*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v98
			*(*int32)(unsafe.Add(mBase, uint32(v57))) = l1
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v99)
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+20)) = uint16(v6)
			*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l3
			switch v6 {
			case 0:
				v109 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+22)) = uint16(v109)
				v112 = base.I32_div_u_s(int32(4096), v55)
				v117 = v112 - v109
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
				v119 = v117
			case 1:
				v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
				v119 = v108
			default:
				v116 = base.I32_div_u_s(int32(65536), v55)
				v117 = v116
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)) = uint16(v117)
				v119 = v117
			}
			v120 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v120)
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+28)) = uint16(v119)
			v123 = int32(65535)
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+26)) = uint16(v123)
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
	var v147 int32
	_ = v147
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v6) <= base.Ui32(v5) {
		v238 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v238
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = v8 + v5
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v10 == int32(0) {
		v238 = v6
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
	v238 = v231
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
	v68 = v18<<(uint(int32(12))%32)&int32(61440) | v39&int32(63)<<(uint(int32(6))%32)
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
	v68 = v18<<(uint(int32(18))%32)&int32(1835008) | v55&v56<<(uint(int32(12))%32) | v61&v56<<(uint(int32(6))%32)
	v69 = int32(3)
	goto L8
L16:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v192 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L17:
	;
	v191 = v184
	goto L16
L18:
	;
	v184 = base.B2i32(v173&int32(255) == int32(9))
	goto L17
L19:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76<<(uint(int32(1))%32))+uint32(_consts[645]))))
	v173 = v166
	goto L18
L20:
	;
	v191 = base.B2i32(base.Ui32(v76-int32(48)) < base.Ui32(int32(10)))
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
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76<<(uint(v142)%32))+uint32(_consts[638]))))
	if v147&v142 != 0 {
		v184 = v142
		goto L17
	} else {
		goto L44
	}
L24:
	;
	v93 = base.I32_div_s(v87+v88, int32(2))
	v95 = v93 << (uint(int32(3)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+uint32(_consts[639])))
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+uint32(_consts[640])))
	if base.Ui32(v104) <= base.Ui32(v76) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v191 = int32(1)
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
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_consts[644]))))
	v173 = v141
	goto L18
L36:
	;
	v122 = base.I32_div_s(v116+v117, int32(2))
	v124 = v122 * int32(12)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_consts[646])))
	if base.Ui32(v127) < base.Ui32(v76) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v173 = int32(0)
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_consts[647])))
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
		v227 = int32(1)
		goto L57
	} else {
		goto L58
	}
L47:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v195 == v191 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v197)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v191)
	if base.Ui32(v76) < base.Ui32(int32(128)) {
		v212 = v197
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v212 + v199
	return v199
L52:
	;
	if base.Ui32(v76) < base.Ui32(int32(2048)) {
		v212 = int32(2)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v76) < base.Ui32(int32(65536)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v211 = int32(3)
	goto L56
L55:
	;
	v211 = int32(4)
	goto L56
L56:
	;
	v212 = v211
	goto L51
L57:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = v227 + v228
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v231) <= base.Ui32(v229) {
		v238 = v231
		goto L1
	} else {
		goto L63
	}
L58:
	;
	if base.Ui32(v76) < base.Ui32(int32(2048)) {
		v227 = int32(2)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v76) < base.Ui32(int32(65536)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v226 = int32(3)
	goto L62
L61:
	;
	v226 = int32(4)
	goto L62
L62:
	;
	v227 = v226
	goto L57
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = v233 + v229
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v235 != 0 {
		v14 = v235
		v16 = v234
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
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
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
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
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
	var v313 int32
	_ = v313
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
	var v358 int32
	_ = v358
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 float64
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 float64
	_ = v431
	var v432 int32
	_ = v432
	var v435 float64
	_ = v435
	var v437 float64
	_ = v437
	var v438 float64
	_ = v438
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 float64
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 float64
	_ = v462
	var v463 int32
	_ = v463
	var v466 float64
	_ = v466
	var v468 float64
	_ = v468
	var v469 float64
	_ = v469
	var v480 int32
	_ = v480
	var v483 float64
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 float32
	_ = v487
	var v489 float64
	_ = v489
	var v491 float64
	_ = v491
	var v496 float64
	_ = v496
	var v501 float64
	_ = v501
	var v504 float64
	_ = v504
	var v505 float32
	_ = v505
	var v507 float64
	_ = v507
	var v509 float64
	_ = v509
	var v514 float64
	_ = v514
	var v519 float64
	_ = v519
	var v524 int32
	_ = v524
	var v527 float64
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 float32
	_ = v531
	var v533 float64
	_ = v533
	var v535 float64
	_ = v535
	var v540 float64
	_ = v540
	var v545 float64
	_ = v545
	var v548 float64
	_ = v548
	var v549 float32
	_ = v549
	var v551 float64
	_ = v551
	var v553 float64
	_ = v553
	var v558 float64
	_ = v558
	var v563 float64
	_ = v563
	var v568 float64
	_ = v568
	var v569 float64
	_ = v569
	var v575 float64
	_ = v575
	var v576 float64
	_ = v576
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 float64
	_ = v639
	var v641 float64
	_ = v641
	var v643 float64
	_ = v643
	var v645 float64
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v663 int32
	_ = v663
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 float64
	_ = v759
	var v762 int32
	_ = v762
	var v764 float64
	_ = v764
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v773 float64
	_ = v773
	var v776 int32
	_ = v776
	var v778 float64
	_ = v778
	var v781 int32
	_ = v781
	var v782 float64
	_ = v782
	var v784 float64
	_ = v784
	var v820 float64
	_ = v820
	var v822 float64
	_ = v822
	var v824 float64
	_ = v824
	var v826 float64
	_ = v826
	var v833 float64
	_ = v833
	var v834 float64
	_ = v834
	var v838 float64
	_ = v838
	var v840 float64
	_ = v840
	var v841 float64
	_ = v841
	var v850 float64
	_ = v850
	var v854 float64
	_ = v854
	var v856 float64
	_ = v856
	var v857 float64
	_ = v857
	var v865 float64
	_ = v865
	var v869 float64
	_ = v869
	var v870 float64
	_ = v870
	var v871 float64
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v883 float64
	_ = v883
	var v884 float64
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 float64
	_ = v892
	var v893 int32
	_ = v893
	var v894 float64
	_ = v894
	var v895 float64
	_ = v895
	var v898 float64
	_ = v898
	var v900 float64
	_ = v900
	var v904 float64
	_ = v904
	var v905 float64
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v915 float64
	_ = v915
	var v917 int64
	_ = v917
	var v919 int64
	_ = v919
	var v920 float64
	_ = v920
	var v922 float64
	_ = v922
	var v928 float64
	_ = v928
	var v930 float64
	_ = v930
	var v933 int32
	_ = v933
	var v935 int64
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 float64
	_ = v943
	var v945 float64
	_ = v945
	var v946 float64
	_ = v946
	var v950 float64
	_ = v950
	var v953 float64
	_ = v953
	var v957 float64
	_ = v957
	var v964 float64
	_ = v964
	var v965 float64
	_ = v965
	var v967 float64
	_ = v967
	var v971 float64
	_ = v971
	var v976 float64
	_ = v976
	var v978 float64
	_ = v978
	var v981 int32
	_ = v981
	var v985 float64
	_ = v985
	var v988 int32
	_ = v988
	var v990 float64
	_ = v990
	var v997 float64
	_ = v997
	var v998 float64
	_ = v998
	var v1003 int32
	_ = v1003
	var v1004 float64
	_ = v1004
	var v1005 float64
	_ = v1005
	var v1006 float64
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1019 float64
	_ = v1019
	var v1021 float64
	_ = v1021
	var v1022 float64
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 float64
	_ = v1024
	var v1027 float64
	_ = v1027
	var v1029 float64
	_ = v1029
	var v1033 float64
	_ = v1033
	var v1034 float64
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1044 float64
	_ = v1044
	var v1046 int64
	_ = v1046
	var v1048 int64
	_ = v1048
	var v1049 float64
	_ = v1049
	var v1051 float64
	_ = v1051
	var v1057 float64
	_ = v1057
	var v1059 float64
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1064 int64
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 float64
	_ = v1072
	var v1074 float64
	_ = v1074
	var v1075 float64
	_ = v1075
	var v1079 float64
	_ = v1079
	var v1082 float64
	_ = v1082
	var v1086 float64
	_ = v1086
	var v1093 float64
	_ = v1093
	var v1094 float64
	_ = v1094
	var v1096 float64
	_ = v1096
	var v1100 float64
	_ = v1100
	var v1105 float64
	_ = v1105
	var v1106 float64
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1113 float64
	_ = v1113
	var v1116 float64
	_ = v1116
	var v1120 float64
	_ = v1120
	var v1121 float64
	_ = v1121
	var v1122 float64
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1131 float64
	_ = v1131
	var v1134 float64
	_ = v1134
	var v1142 float64
	_ = v1142
	var v1148 float64
	_ = v1148
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	v29 = float64(0)
	v46 = m.G0
	v48 = v46 - int32(96)
	m.G0 = v48
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v52 = float64(1)
	if l2 == int32(2) {
		v820 = v52
		v822 = v52
		v824 = v29
		v826 = v29
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L33
	} else {
		goto L255
	}
L2:
	;
	if base.F64_le(v50, float64(0)) != 0 {
		goto L188
	} else {
		goto L189
	}
L3:
	;
	if l3 == int32(0) {
		v820 = v52
		v822 = v52
		v824 = v29
		v826 = v29
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if l6 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = l6
	goto L7
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v59 = v58
	goto L7
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if l7 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v64 = l7
	goto L10
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l5)+64))
	v64 = v63
	goto L10
L10:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v62 != v67 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v70 != v72 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if v74 != v75 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+16)))
	if v77 != v78 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+116))
	if v82 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	v703 = int32(0)
	if v700 == v703 {
		goto L156
	} else {
		goto L157
	}
L16:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v196 = m.G0
	v198 = v196 - int32(96)
	m.G0 = v198
	v201 = v48 + int32(80)
	v202 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v202
	v207 = v48 + int32(72)
	v208 = int64(4607182418800017408)
	*(*int64)(unsafe.Add(mBase, uint32(v207))) = v208
	v211 = v48 + int32(88)
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = v208
	if v195 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v85 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v93 = int32(0)
	goto L19
L19:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v88+v93<<(uint(int32(2))%32))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v139 != v62 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L16
L21:
	;
	v148 = v93 + int32(1)
	if v85 != v148 {
		v93 = v148
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v141 != v70 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v143 != v74 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+12)))
	if v145 == v77 {
		v663 = v138
		goto L15
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	goto L20
L27:
	;
	m.G0 = v198 + int32(96)
	v622 = int32(4554240)
	v623 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v625
	v628 = F_palloc(m, int32(48))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L33
	} else {
		goto L153
	}
L28:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	if v216 != int32(17) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v195)+28))
	if v219 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v222 < int32(2) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v226 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v195)+24))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	F_examine_variable(m, l0, v231, int32(0), v198-int32(-64))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return
L34:
	;
	F_examine_variable(m, l0, v226, int32(0), v198+int32(32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v242 = F_get_opfamily_method(m, v62)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	F_get_op_opfamily_properties(m, v230, v62, int32(0), v198+int32(28), v198+int32(24), v198+int32(20))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	switch v74 - int32(1) {
	case 0:
		goto L41
	default:
		goto L38
	case 4:
		goto L40
	}
L38:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v198)+72))
	if v596 != 0 {
		goto L147
	} else {
		goto L148
	}
L39:
	;
	if v354 == int32(0) {
		goto L38
	} else {
		goto L72
	}
L40:
	;
	v296 = int32(1)
	v299 = F_IndexAmTranslateCompareType(m, v296, v242, v62, v296)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L33
	} else {
		goto L55
	}
L41:
	;
	v255 = int32(1)
	v257 = F_IndexAmTranslateCompareType(m, v255, v242, v62, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v261 = F_IndexAmTranslateCompareType(m, int32(2), v242, v62, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v263 = base.I32_extend16_s(v261)
	v264 = base.I32_extend16_s(v257)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	if v265 == v266 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v269 = F_get_opfamily_member(m, v62, v265, v265, v264)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L33
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v276 = F_get_opfamily_member(m, v62, v265, v266, v264)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L33
	} else {
		goto L49
	}
L47:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v273 = F_get_opfamily_member(m, v62, v271, v272, v263)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L33
	} else {
		goto L48
	}
L48:
	;
	v350 = v269
	v351 = v269
	v352 = v269
	v353 = v269
	v354 = v269
	v355 = int32(0)
	v356 = v273
	v357 = v269
	v358 = v273
	goto L39
L49:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v280 = F_get_opfamily_member(m, v62, v278, v279, v263)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L33
	} else {
		goto L50
	}
L50:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v283 = F_get_opfamily_member(m, v62, v282, v282, v264)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L33
	} else {
		goto L51
	}
L51:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v286 = F_get_opfamily_member(m, v62, v285, v285, v264)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L33
	} else {
		goto L52
	}
L52:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v290 = F_get_opfamily_member(m, v62, v288, v289, v264)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L33
	} else {
		goto L53
	}
L53:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v294 = F_get_opfamily_member(m, v62, v292, v293, v263)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	v350 = v286
	v351 = v276
	v352 = v283
	v353 = v290
	v354 = v283
	v355 = int32(0)
	v356 = v294
	v357 = v286
	v358 = v280
	goto L39
L55:
	;
	v303 = F_IndexAmTranslateCompareType(m, int32(5), v242, v62, int32(1))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L33
	} else {
		goto L56
	}
L56:
	;
	v305 = base.I32_extend16_s(v303)
	v308 = F_IndexAmTranslateCompareType(m, int32(4), v242, v62, int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L33
	} else {
		goto L57
	}
L57:
	;
	v310 = base.I32_extend16_s(v308)
	v311 = base.I32_extend16_s(v299)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	if v312 == v313 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v315 = F_get_opfamily_member(m, v62, v312, v312, v305)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L33
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v324 = F_get_opfamily_member(m, v62, v312, v313, v305)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L33
	} else {
		goto L64
	}
L61:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v319 = F_get_opfamily_member(m, v62, v317, v318, v310)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v322 = F_get_opfamily_member(m, v62, v321, v321, v311)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L33
	} else {
		goto L63
	}
L63:
	;
	v350 = v322
	v351 = v315
	v352 = v322
	v353 = v315
	v354 = v315
	v355 = v296
	v356 = v319
	v357 = v315
	v358 = v319
	goto L39
L64:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v328 = F_get_opfamily_member(m, v62, v326, v327, v310)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L33
	} else {
		goto L65
	}
L65:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v331 = F_get_opfamily_member(m, v62, v330, v330, v305)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L33
	} else {
		goto L66
	}
L66:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v334 = F_get_opfamily_member(m, v62, v333, v333, v305)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L33
	} else {
		goto L67
	}
L67:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v337 = F_get_opfamily_member(m, v62, v336, v336, v311)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L33
	} else {
		goto L68
	}
L68:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v340 = F_get_opfamily_member(m, v62, v339, v339, v311)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L33
	} else {
		goto L69
	}
L69:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v344 = F_get_opfamily_member(m, v62, v342, v343, v305)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L33
	} else {
		goto L70
	}
L70:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v348 = F_get_opfamily_member(m, v62, v346, v347, v310)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L33
	} else {
		goto L71
	}
L71:
	;
	v350 = v340
	v351 = v324
	v352 = v337
	v353 = v344
	v354 = v331
	v355 = v296
	v356 = v348
	v357 = v334
	v358 = v328
	goto L39
L72:
	;
	if v357 == int32(0) {
		goto L38
	} else {
		goto L73
	}
L73:
	;
	if v352 == int32(0) {
		goto L38
	} else {
		goto L74
	}
L74:
	;
	if v350 == int32(0) {
		goto L38
	} else {
		goto L75
	}
L75:
	;
	if v351 == int32(0) {
		goto L38
	} else {
		goto L76
	}
L76:
	;
	if v358 == int32(0) {
		goto L38
	} else {
		goto L77
	}
L77:
	;
	if v353 == int32(0) {
		goto L38
	} else {
		goto L78
	}
L78:
	;
	if v356 == int32(0) {
		goto L38
	} else {
		goto L79
	}
L79:
	;
	if v355 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v421 = F_scalarineqsel(m, l0, v358, v355, int32(1), v229, v198-int32(-64), v419, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L33
	} else {
		goto L92
	}
L81:
	;
	v384 = F_get_variable_range(m, v198-int32(-64), v352, v229, v198+int32(16), v198+int32(12))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L33
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v402 = F_get_variable_range(m, v198-int32(-64), v352, v229, v198+int32(12), v198+int32(16))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L33
	} else {
		goto L88
	}
L84:
	;
	if v384 == int32(0) {
		goto L38
	} else {
		goto L85
	}
L85:
	;
	v394 = F_get_variable_range(m, v198+int32(32), v350, v229, v198+int32(8), v198+int32(4))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L33
	} else {
		goto L86
	}
L86:
	;
	if v394 != 0 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	goto L38
L88:
	;
	if v402 == int32(0) {
		goto L38
	} else {
		goto L89
	}
L89:
	;
	v412 = F_get_variable_range(m, v198+int32(32), v350, v229, v198+int32(4), v198+int32(8))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L33
	} else {
		goto L90
	}
L90:
	;
	if v412 == int32(0) {
		goto L38
	} else {
		goto L91
	}
L91:
	;
	goto L80
L92:
	;
	if base.F64_ne(v421, float64(0.3333333333333333)) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v211))) = v421
	goto L95
L94:
	;
	goto L95
L95:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v431 = F_scalarineqsel(m, l0, v356, v355, int32(1), v229, v198+int32(32), v429, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L33
	} else {
		goto L97
	}
L96:
	;
	v438 = *(*float64)(unsafe.Add(mBase, uint32(v211)))
	if base.F64_gt(v438, v437) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	if base.F64_eq(v431, float64(0.3333333333333333)) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v435 = *(*float64)(unsafe.Add(mBase, uint32(v207)))
	v437 = v435
	goto L96
L99:
	;
	goto L100
L100:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v207))) = v431
	v437 = v431
	goto L96
L101:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v452 = F_scalarineqsel(m, l0, v351, v355, int32(0), v229, v198-int32(-64), v450, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L33
	} else {
		goto L106
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v207))) = int64(4607182418800017408)
	if base.F64_gt(v437, v438) != 0 {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = int64(4607182418800017408)
	goto L101
L105:
	;
	goto L104
L106:
	;
	if base.F64_ne(v452, float64(0.3333333333333333)) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v452
	goto L109
L108:
	;
	goto L109
L109:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v462 = F_scalarineqsel(m, l0, v353, v355, int32(0), v229, v198+int32(32), v460, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L33
	} else {
		goto L111
	}
L110:
	;
	v469 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	if base.F64_lt(v469, v468) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L111:
	;
	if base.F64_eq(v462, float64(0.3333333333333333)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v201)))
	v468 = v466
	goto L110
L113:
	;
	goto L114
L114:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v201))) = v462
	v468 = v462
	goto L110
L115:
	;
	if v77 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = int64(0)
	if base.F64_lt(v468, v469) != 0 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(0)
	goto L115
L119:
	;
	goto L118
L120:
	;
	v568 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v569 = *(*float64)(unsafe.Add(mBase, uint32(v211)))
	if base.F64_ge(v568, v569) != 0 {
		goto L143
	} else {
		goto L144
	}
L121:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v198)+72))
	if v480 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v198)+40))
	if v524 == int32(0) {
		goto L120
	} else {
		goto L133
	}
L123:
	;
	v483 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480)+16))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+22)))
	v486 = v484 + v485
	v487 = *(*float32)(unsafe.Add(mBase, uint32(v486)+8))
	v489 = base.F64_add(v483, base.F64_promote_f32(v487))
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v489
	v491 = float64(0)
	if base.F64_lt(v489, v491) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v504 = *(*float64)(unsafe.Add(mBase, uint32(v211)))
	v505 = *(*float32)(unsafe.Add(mBase, uint32(v486)+8))
	v507 = base.F64_add(v504, base.F64_promote_f32(v505))
	*(*float64)(unsafe.Add(mBase, uint32(v211))) = v507
	v509 = float64(0)
	if base.F64_lt(v507, v509) == int32(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v496 = float64(1)
	if base.F64_gt(v489, v496) == int32(0) {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	v501 = v491
	goto L127
L127:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v501
	goto L124
L128:
	;
	v501 = v496
	goto L127
L129:
	;
	v514 = float64(1)
	if base.F64_gt(v507, v514) == int32(0) {
		goto L122
	} else {
		goto L132
	}
L130:
	;
	v519 = v509
	goto L131
L131:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v211))) = v519
	goto L122
L132:
	;
	v519 = v514
	goto L131
L133:
	;
	v527 = *(*float64)(unsafe.Add(mBase, uint32(v201)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+22)))
	v530 = v528 + v529
	v531 = *(*float32)(unsafe.Add(mBase, uint32(v530)+8))
	v533 = base.F64_add(v527, base.F64_promote_f32(v531))
	*(*float64)(unsafe.Add(mBase, uint32(v201))) = v533
	v535 = float64(0)
	if base.F64_lt(v533, v535) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v548 = *(*float64)(unsafe.Add(mBase, uint32(v207)))
	v549 = *(*float32)(unsafe.Add(mBase, uint32(v530)+8))
	v551 = base.F64_add(v548, base.F64_promote_f32(v549))
	*(*float64)(unsafe.Add(mBase, uint32(v207))) = v551
	v553 = float64(0)
	if base.F64_lt(v551, v553) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v540 = float64(1)
	if base.F64_gt(v533, v540) == int32(0) {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	v545 = v535
	goto L137
L137:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v201))) = v545
	goto L134
L138:
	;
	v545 = v540
	goto L137
L139:
	;
	v558 = float64(1)
	if base.F64_gt(v551, v558) == int32(0) {
		goto L120
	} else {
		goto L142
	}
L140:
	;
	v563 = v553
	goto L141
L141:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v207))) = v563
	goto L120
L142:
	;
	v563 = v558
	goto L141
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = int64(4607182418800017408)
	goto L145
L144:
	;
	goto L145
L145:
	;
	v575 = *(*float64)(unsafe.Add(mBase, uint32(v201)))
	v576 = *(*float64)(unsafe.Add(mBase, uint32(v207)))
	if base.F64_ge(v575, v576) == int32(0) {
		goto L38
	} else {
		goto L146
	}
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v207))) = int64(4607182418800017408)
	goto L38
L147:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v198)+76))
	m.T0[v597].(func(*base.Module, int32))(m, v596)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L33
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v198)+40))
	if v600 == int32(0) {
		goto L27
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v198)+44))
	m.T0[v603].(func(*base.Module, int32))(m, v600)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L33
	} else {
		goto L152
	}
L152:
	;
	goto L27
L153:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v628)+4)) = v633
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v628)+8)) = v635
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v628)+12)) = uint8(v637)
	v639 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	*(*float64)(unsafe.Add(mBase, uint32(v628)+16)) = v639
	v641 = *(*float64)(unsafe.Add(mBase, uint32(v48)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v628)+24)) = v641
	v643 = *(*float64)(unsafe.Add(mBase, uint32(v48)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v628)+32)) = v643
	v645 = *(*float64)(unsafe.Add(mBase, uint32(v48)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v628)+40)) = v645
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v81)+116))
	v648 = F_lappend(m, v647, v628)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L33
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+116)) = v648
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v623
	v663 = v628
	goto L15
L155:
	;
	if v756 != 0 {
		goto L169
	} else {
		goto L170
	}
L156:
	;
	v756 = int32(1)
	goto L155
L157:
	;
	goto L158
L158:
	;
	if v702 == int32(0) {
		v747 = v703
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v756 = v747
	goto L155
L160:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	if v713 < v712 {
		v747 = v703
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v715 = int32(1)
	if v712 <= v715 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v718 = v715
	goto L164
L163:
	;
	v718 = v712
	goto L164
L164:
	;
	v719 = int32(8)
	v724 = int32(0)
	goto L165
L165:
	;
	v731 = v724 << (uint(int32(2)) % 32)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v700+v719+v731)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v731+(v702+v719))))
	v738 = v733 & (v735 ^ int32(-1))
	v740 = base.B2i32(v738 == int32(0))
	if v738 != 0 {
		v747 = v740
		goto L159
	} else {
		goto L167
	}
L166:
	;
	v747 = v740
	goto L159
L167:
	;
	v742 = v724 + int32(1)
	if v742 != v718 {
		v724 = v742
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v757 = int32(32)
	goto L171
L170:
	;
	v757 = int32(16)
	goto L171
L171:
	;
	v759 = *(*float64)(unsafe.Add(mBase, uint32(v663+v757)))
	if v756 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v762 = int32(40)
	goto L174
L173:
	;
	v762 = int32(24)
	goto L174
L174:
	;
	v764 = *(*float64)(unsafe.Add(mBase, uint32(v663+v762)))
	v766 = l2 & int32(-5)
	if v766 == int32(1) {
		v820 = v52
		v822 = v764
		v824 = v759
		v826 = v29
		goto L2
	} else {
		goto L175
	}
L175:
	;
	if v756 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v771 = int32(24)
	goto L178
L177:
	;
	v771 = int32(40)
	goto L178
L178:
	;
	v773 = *(*float64)(unsafe.Add(mBase, uint32(v663+v771)))
	if v756 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v776 = int32(16)
	goto L181
L180:
	;
	v776 = int32(32)
	goto L181
L181:
	;
	v778 = *(*float64)(unsafe.Add(mBase, uint32(v663+v776)))
	v781 = base.B2i32(v766 == int32(3))
	if v766 == int32(3) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v782 = float64(1)
	goto L184
L183:
	;
	v782 = v764
	goto L184
L184:
	;
	if v766 == int32(3) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v784 = float64(0)
	goto L187
L186:
	;
	v784 = v759
	goto L187
L187:
	;
	v820 = v773
	v822 = v782
	v824 = v784
	v826 = v778
	goto L2
L188:
	;
	v833 = float64(1)
	goto L190
L189:
	;
	v833 = v50
	goto L190
L190:
	;
	v834 = float64(1e+100)
	if base.F64_le(v51, float64(0)) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v838 = float64(1)
	goto L193
L192:
	;
	v838 = v51
	goto L193
L193:
	;
	v840 = float64(1e+100)
	v841 = base.F64_mul(v838, v820)
	if base.F64_gt(v841, v840) != 0 {
		v854 = v840
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v856 = base.F64_nearest(base.F64_mul(v838, v826))
	v857 = base.F64_mul(v833, v822)
	if base.F64_gt(v857, float64(1e+100)) != 0 {
		v869 = v834
		goto L198
	} else {
		goto L199
	}
L195:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v841)&int64(9223372036854775807)) {
		v854 = float64(1e+100)
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v850 = float64(1)
	if base.F64_le(v841, v850) != 0 {
		v854 = v850
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v854 = base.F64_nearest(v841)
	goto L194
L198:
	;
	v870 = base.F64_nearest(base.F64_mul(v833, v824))
	v871 = base.F64_div(v856, v838)
	v873 = int32(*(*uint8)(unsafe.Add(mBase, _consts[606])))
	if l6 != 0 {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v857)&int64(9223372036854775807)) {
		v869 = v834
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v865 = float64(1)
	if base.F64_le(v857, v865) != 0 {
		v869 = v865
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v869 = base.F64_nearest(v857)
	goto L198
L202:
	;
	v1022 = base.F64_div(v870, v833)
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	if l7 != 0 {
		goto L232
	} else {
		goto L233
	}
L203:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	if l8 <= int32(0) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	goto L205
L205:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	v1004 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v1005 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v1006 = base.F64_sub(v1004, v1005)
	v1012 = v1003
	v1019 = v1006
	v1021 = base.F64_add(base.F64_mul(v1006, v871), base.F64_add(v1005, float64(0)))
	goto L202
L206:
	;
	v998 = base.F64_sub(v997, v990)
	v1012 = v988
	v1019 = v998
	v1021 = base.F64_add(base.F64_mul(v998, v871), base.F64_add(v990, float64(0)))
	goto L202
L207:
	;
	v895 = float64(2)
	if base.F64_lt(v838, v895) != 0 {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, _consts[607])))
	if v878&int32(1) == int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v883 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v884 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)+32))
	v888 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	F_cost_incremental_sort(m, v48, l0, l6, l8, v874, v883, v884, v838, v886, v888, float64(-1))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L33
	} else {
		goto L210
	}
L210:
	;
	v892 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	v894 = *(*float64)(unsafe.Add(mBase, uint32(v48)+56))
	v988 = v893
	v990 = v892
	v997 = v894
	goto L206
L211:
	;
	v898 = v895
	goto L213
L212:
	;
	v898 = v838
	goto L213
L213:
	;
	v900 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v904 = base.F64_mul(v898, base.F64_add(base.F64_add(v900, v900), float64(0)))
	v905 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)+32))
	v915 = base.F64_mul(v838, base.F64_convert_i32_u((v907+int32(7))&int32(-8)+int32(24)))
	v917 = int64(*(*int32)(unsafe.Add(mBase, _consts[45])))
	v919 = v917 << (uint(int64(10)) % 64)
	v920 = base.F64_convert_i64_s(v919)
	if base.F64_gt(v915, v920) != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, _consts[608])))
	v985 = base.F64_add(v905, v976)
	v988 = v874 + (v981 ^ int32(1))
	v990 = v985
	v997 = base.F64_add(v985, base.F64_mul(v898, v978))
	goto L206
L215:
	;
	v922 = F_log(m, v898)
	mBase = m.M
	v928 = base.F64_ceil(base.F64_mul(v915, float64(0.0001220703125)))
	v930 = base.F64_div(v915, v920)
	v933 = int32(6)
	v935 = base.I64_div_s(v919, int64(278528))
	v936 = base.I32_wrap_i64(v935)
	if v936 <= v933 {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	goto L217
L217:
	;
	v965 = base.F64_add(v898, v898)
	if base.F64_lt(v965, v898) != 0 {
		goto L228
	} else {
		goto L229
	}
L218:
	;
	v943 = base.F64_convert_i32_s(v942)
	if base.F64_gt(v930, v943) != 0 {
		goto L225
	} else {
		goto L226
	}
L219:
	;
	v939 = v933
	goto L221
L220:
	;
	v939 = v936
	goto L221
L221:
	;
	if int32(500) <= v939 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v942 = int32(500)
	goto L224
L223:
	;
	v942 = v939
	goto L224
L224:
	;
	goto L218
L225:
	;
	v945 = F_log(m, v930)
	mBase = m.M
	v946 = F_log(m, v943)
	mBase = m.M
	v950 = base.F64_ceil(base.F64_div(v945, v946))
	goto L227
L226:
	;
	v950 = float64(1)
	goto L227
L227:
	;
	v953 = *(*float64)(unsafe.Add(mBase, _consts[604]))
	v957 = *(*float64)(unsafe.Add(mBase, _consts[603]))
	v964 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v976 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v928, v928), v950), base.F64_add(base.F64_mul(v953, float64(0.75)), base.F64_mul(v957, float64(0.25)))), base.F64_mul(base.F64_div(v922, float64(0.693147180559945)), v904))
	v978 = v964
	goto L214
L228:
	;
	v967 = F_log(m, v965)
	mBase = m.M
	v976 = base.F64_mul(base.F64_div(v967, float64(0.693147180559945)), v904)
	v978 = v900
	goto L214
L229:
	;
	goto L230
L230:
	;
	v971 = F_log(m, v898)
	mBase = m.M
	v976 = base.F64_mul(base.F64_div(v971, float64(0.693147180559945)), v904)
	v978 = v900
	goto L214
L231:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+72)) = v870
	*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = v856
	*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v869
	*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = v854
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1134
	v1142 = base.F64_mul(base.F64_sub(base.F64_div(v869, v833), v1022), v1131)
	*(*float64)(unsafe.Add(mBase, uint32(l1)+32)) = v1142
	v1148 = base.F64_add(base.F64_mul(v1019, base.F64_sub(base.F64_div(v854, v838), v871)), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v1148
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1012 + (v873^int32(1))&int32(255) + v1126
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v1142, base.F64_add(v1148, v1134))
	m.G0 = v48 + int32(96)
	return
L232:
	;
	v1024 = float64(2)
	if base.F64_lt(v833, v1024) != 0 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	v1120 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v1121 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v1122 = base.F64_sub(v1120, v1121)
	v1126 = v1023
	v1131 = v1122
	v1134 = base.F64_add(base.F64_mul(v1122, v1022), base.F64_add(v1021, v1121))
	goto L231
L235:
	;
	v1027 = v1024
	goto L237
L236:
	;
	v1027 = v833
	goto L237
L237:
	;
	v1029 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v1033 = base.F64_mul(v1027, base.F64_add(base.F64_add(v1029, v1029), float64(0)))
	v1034 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+32))
	v1044 = base.F64_mul(v833, base.F64_convert_i32_u((v1036+int32(7))&int32(-8)+int32(24)))
	v1046 = int64(*(*int32)(unsafe.Add(mBase, _consts[45])))
	v1048 = v1046 << (uint(int64(10)) % 64)
	v1049 = base.F64_convert_i64_s(v1048)
	if base.F64_gt(v1044, v1049) != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[608])))
	v1113 = base.F64_add(v1034, v1105)
	v1116 = base.F64_sub(base.F64_add(v1113, base.F64_mul(v1027, v1106)), v1113)
	v1126 = v1023 + (v1109 ^ int32(1))
	v1131 = v1116
	v1134 = base.F64_add(base.F64_mul(v1116, v1022), base.F64_add(v1021, v1113))
	goto L231
L239:
	;
	v1051 = F_log(m, v1027)
	mBase = m.M
	v1057 = base.F64_ceil(base.F64_mul(v1044, float64(0.0001220703125)))
	v1059 = base.F64_div(v1044, v1049)
	v1062 = int32(6)
	v1064 = base.I64_div_s(v1048, int64(278528))
	v1065 = base.I32_wrap_i64(v1064)
	if v1065 <= v1062 {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	goto L241
L241:
	;
	v1094 = base.F64_add(v1027, v1027)
	if base.F64_lt(v1094, v1027) != 0 {
		goto L252
	} else {
		goto L253
	}
L242:
	;
	v1072 = base.F64_convert_i32_s(v1071)
	if base.F64_gt(v1059, v1072) != 0 {
		goto L249
	} else {
		goto L250
	}
L243:
	;
	v1068 = v1062
	goto L245
L244:
	;
	v1068 = v1065
	goto L245
L245:
	;
	if int32(500) <= v1068 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1071 = int32(500)
	goto L248
L247:
	;
	v1071 = v1068
	goto L248
L248:
	;
	goto L242
L249:
	;
	v1074 = F_log(m, v1059)
	mBase = m.M
	v1075 = F_log(m, v1072)
	mBase = m.M
	v1079 = base.F64_ceil(base.F64_div(v1074, v1075))
	goto L251
L250:
	;
	v1079 = float64(1)
	goto L251
L251:
	;
	v1082 = *(*float64)(unsafe.Add(mBase, _consts[604]))
	v1086 = *(*float64)(unsafe.Add(mBase, _consts[603]))
	v1093 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v1105 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v1057, v1057), v1079), base.F64_add(base.F64_mul(v1082, float64(0.75)), base.F64_mul(v1086, float64(0.25)))), base.F64_mul(base.F64_div(v1051, float64(0.693147180559945)), v1033))
	v1106 = v1093
	goto L238
L252:
	;
	v1096 = F_log(m, v1094)
	mBase = m.M
	v1105 = base.F64_mul(base.F64_div(v1096, float64(0.693147180559945)), v1033)
	v1106 = v1029
	goto L238
L253:
	;
	goto L254
L254:
	;
	v1100 = F_log(m, v1027)
	mBase = m.M
	v1105 = base.F64_mul(base.F64_div(v1100, float64(0.693147180559945)), v1033)
	v1106 = v1029
	goto L238
L255:
	;
	F_errmsg_internal(m, int32(288375), int32(0))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L33
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(521926), int32(3615), int32(288335))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L33
	} else {
		goto L257
	}
L257:
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 float64
	_ = v158
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 float64
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 float64
	_ = v195
	var v206 int32
	_ = v206
	var v215 float64
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 float64
	_ = v220
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v226 float64
	_ = v226
	var v229 float64
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v251 float64
	_ = v251
	var v258 float64
	_ = v258
	var v261 float64
	_ = v261
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v278 float64
	_ = v278
	var v279 float64
	_ = v279
	var v295 float64
	_ = v295
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v301 float64
	_ = v301
	var v305 float64
	_ = v305
	var v306 float64
	_ = v306
	var v307 float64
	_ = v307
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v314 int32
	_ = v314
	var v319 float64
	_ = v319
	var v326 float64
	_ = v326
	var v330 float64
	_ = v330
	v7 = float64(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[605])))
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
	v295 = base.F64_add(v32, float64(-1))
	v297 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v298 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v301 = base.F64_add(base.F64_sub(v297, v298), float64(0))
	if base.F64_gt(v32, float64(1)) != 0 {
		goto L48
	} else {
		goto L49
	}
L2:
	;
	v269 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v270 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v278 = v269
	v279 = v270
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
	v112 = *(*float64)(unsafe.Add(mBase, _consts[434]))
	v114 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v118 = base.F64_mul(base.F64_mul(v112, base.F64_convert_i32_s(v114)), float64(1024))
	v119 = float64(4.294967295e+09)
	if base.F64_lt(v118, v119) != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v75 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v77 = base.F64_mul(v75, v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+32))
	v87 = base.F64_mul(v76, base.F64_convert_i32_u((v79+int32(7))&int32(-8)+int32(24)))
	v89 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if base.F64_gt(v87, base.F64_convert_i32_u(v89<<(uint(int32(10))%32))) == int32(0) {
		v278 = v77
		v279 = v7
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v46 = *(*float64)(unsafe.Add(mBase, _consts[600]))
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v48 = base.F64_mul(v46, v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
	v58 = base.F64_mul(v47, base.F64_convert_i32_u((v50+int32(7))&int32(-8)+int32(24)))
	v60 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if base.F64_gt(v58, base.F64_convert_i32_u(v60<<(uint(int32(10))%32))) == int32(0) {
		v278 = v48
		v279 = v7
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
	v278 = base.F64_sub(v36, v37)
	v279 = v7
	goto L1
L8:
	;
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v278 = base.F64_sub(v42, v43)
	v279 = v7
	goto L1
L9:
	;
	v68 = *(*float64)(unsafe.Add(mBase, _consts[604]))
	v278 = base.F64_add(base.F64_mul(v68, base.F64_ceil(base.F64_mul(v58, float64(0.0001220703125)))), v48)
	v279 = v7
	goto L1
L10:
	;
	v97 = *(*float64)(unsafe.Add(mBase, _consts[604]))
	v278 = base.F64_add(base.F64_mul(v97, base.F64_ceil(base.F64_mul(v87, float64(0.0001220703125)))), v77)
	v279 = v7
	goto L1
L11:
	;
	v143 = base.F64_add(base.F64_add(base.F64_mul(v106, float64(8)), float64(28)), base.F64_mul(v106, base.F64_convert_i32_u((v105+int32(7))&int32(-8)+int32(24))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	if v144 == int32(0) {
		goto L19
	} else {
		goto L20
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
	if base.F64_lt(v122, float64(4.294967296e+09))&base.F64_ge(v122, float64(0)) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v128 = base.I32_trunc_f64_u(v122)
	v130 = v128
	goto L11
L16:
	;
	goto L17
L17:
	;
	v130 = int32(0)
	goto L11
L18:
	;
	v215 = F_estimate_num_groups(m, l0, v206, v107, int32(0), v26+int32(12))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L25
	} else {
		goto L29
	}
L19:
	;
	v195 = v143
	v206 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v148 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v149 <= v148 {
		v195 = v143
		v206 = v144
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v158 = v143
	v168 = v148
	goto L23
L23:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v168<<(uint(int32(2))%32))))
	v180 = F_get_expr_width(m, l0, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v195 = v183
	v206 = v188
	goto L18
L25:
	;
	return
L26:
	;
	v183 = base.F64_add(v158, base.F64_convert_i32_s(v180))
	v185 = v168 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v185 < v186 {
		v158 = v183
		v168 = v185
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+96)) = v237
	v240 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v243 = *(*float64)(unsafe.Add(mBase, _consts[600]))
	v248 = base.F64_lt(v223, v220)
	if v248 != 0 {
		goto L42
	} else {
		goto L43
	}
L29:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v217&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v220 = v107
	goto L32
L31:
	;
	v220 = v215
	goto L32
L32:
	;
	v223 = base.F64_floor(base.F64_div(base.F64_convert_i32_u(v130), v195))
	if base.F64_gt(v223, v220) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v225 = v220
	goto L35
L34:
	;
	v225 = v223
	goto L35
L35:
	;
	v226 = float64(4.294967295e+09)
	if base.F64_lt(v225, v226) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v229 = v225
	goto L38
L37:
	;
	v229 = v226
	goto L38
L38:
	;
	if base.F64_lt(v229, float64(4.294967296e+09))&base.F64_ge(v229, float64(0)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v235 = base.I32_trunc_f64_u(v229)
	v237 = v235
	goto L28
L40:
	;
	goto L41
L41:
	;
	v237 = int32(0)
	goto L28
L42:
	;
	v249 = v223
	goto L44
L43:
	;
	v249 = v220
	goto L44
L44:
	;
	v251 = base.F64_sub(float64(1), base.F64_div(v249, v220))
	if v248 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v258 = v220
	goto L47
L46:
	;
	v258 = v223
	goto L47
L47:
	;
	v261 = base.F64_sub(float64(1), base.F64_mul(base.F64_div(base.F64_sub(v107, v220), v107), base.F64_div(v223, v258)))
	v278 = base.F64_add(base.F64_add(base.F64_mul(v240, v106), v243), base.F64_add(base.F64_mul(base.F64_mul(base.F64_div(v240, float64(10)), v251), v106), base.F64_add(base.F64_mul(v243, v251), base.F64_add(base.F64_mul(v108, v261), v240))))
	v279 = base.F64_add(v243, base.F64_mul(v109, v261))
	goto L1
L48:
	;
	v305 = base.F64_add(base.F64_mul(v295, v279), v301)
	goto L50
L49:
	;
	v305 = v301
	goto L50
L50:
	;
	v306 = base.F64_sub(v278, v279)
	v307 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v308 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v309 = base.F64_sub(v307, v308)
	if l2&int32(-2) != int32(4) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v326
	v330 = base.F64_add(base.F64_add(v298, v308), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v330
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v330, v326)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29 + (v31 ^ int32(1)) + v28
	m.G0 = v26 + int32(16)
	return
L52:
	;
	v319 = base.F64_add(v309, v305)
	if base.F64_gt(v32, float64(1)) == int32(0) {
		v326 = v319
		goto L51
	} else {
		goto L57
	}
L53:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v314 != int32(1) {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = v306
	*(*float64)(unsafe.Add(mBase, uint32(l1)+32)) = v309
	v326 = v305
	goto L51
L56:
	;
	goto L55
L57:
	;
	v326 = base.F64_add(base.F64_mul(v295, v306), v319)
	goto L51
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v2 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		v8 = F_GetSysCacheHashValue(m, int32(21), v6, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1057])) = v8
			F_CacheRegisterSyscacheCallback(m, int32(9), int32(1255), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(11), int32(1255), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(21), int32(1255), int32(0))
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
	var v36 int32
	_ = v36
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v3 = int32(0)
	if l0 == v3 {
		v90 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v90
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
	v88 = F_expression_tree_walker_impl(m, l0, int32(844), l1)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L26
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
		v90 = v3
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v56-v55 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v40 = v31
	v41 = v32
	goto L15
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v55 = v44
	v56 = v45
	goto L12
L17:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v58 != v59 {
		v90 = v3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v62 = l0 + int32(84)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v64 = F_copyObjectImpl(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_IncrementVarSublevelsUp(m, v64, v66, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v72
	v74 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v74
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v72)
	*(*int64)(unsafe.Add(mBase, uint32(v62))) = v74
	return v72
L25:
	;
	goto L24
L26:
	;
	v90 = v88
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
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var __phi242 int32
	_ = __phi242
	var v243 int32
	_ = v243
	var __phi243 int32
	_ = __phi243
	var v244 int32
	_ = v244
	var __phi244 int32
	_ = __phi244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
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
	if v185 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	if v25 == int32(0) {
		v203 = v17
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
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v32<<(uint(int32(2))%32)+(int32(base.Ui32(v23)>>(uint(v36)%32))&int32(2047)+int32(base.Ui32(v23)>>(uint(int32(12))%32))+v36)&int32(4194302)))))
	v185 = v48
	goto L3
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v55 = int32(1)
	v66 = v19 + v51<<(uint(int32(2))%32) + (int32(base.Ui32(v23)>>(uint(v55)%32))&int32(2047)+int32(base.Ui32(v23)>>(uint(int32(12))%32))+v55)&int32(4194302)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66))))
	if v67 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v71 = v67 & int32(3)
	v72 = base.I32_extend8_s(v27)
	if base.Ui32(v67) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v71 == int32(0) {
		v185 = v139
		goto L3
	} else {
		goto L17
	}
L11:
	;
	v137 = v66
	v139 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v82 = v66
	v84 = int32(0)
	v88 = v5
	goto L14
L14:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82+int32(2)))))
	v98 = int32(14)
	v101 = int32(1)
	v104 = int32(4)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82+v104))))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82+int32(6)))))
	v123 = v82 + int32(8)
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v130 = int32(base.Ui32(v72)>>(uint(int32(base.Ui32(v97)>>(uint(v98)%32)))%32))&v101 + v84 + int32(base.Ui32(v72)>>(uint(int32(base.Ui32(v106)>>(uint(v98)%32)))%32))&v101 + int32(base.Ui32(v72)>>(uint(int32(base.Ui32(v115)>>(uint(v98)%32)))%32))&v101 + int32(base.Ui32(v72)>>(uint(int32(base.Ui32(v124)>>(uint(v98)%32)))%32))&v101
	v132 = v88 + v104
	if v132 != v67&int32(65532) {
		v82 = v123
		v84 = v130
		v88 = v132
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v137 = v123
	v139 = v130
	goto L10
L16:
	;
	goto L15
L17:
	;
	v155 = v137
	v157 = v139
	v160 = v5
	goto L18
L18:
	;
	v169 = v155 + int32(2)
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169))))
	v174 = int32(1)
	v176 = int32(base.Ui32(v72)>>(uint(int32(base.Ui32(v170)>>(uint(int32(14))%32)))%32))&v174 + v157
	v178 = v160 + v174
	if v178 != v71 {
		v155 = v169
		v157 = v176
		v160 = v178
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v185 = v176
	goto L3
L20:
	;
	goto L19
L21:
	;
	v203 = v185
	goto L2
L22:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v369) < base.Ui32(v360) {
		goto L71
	} else {
		goto L72
	}
L23:
	;
	v216 = int32(1)
	v217 = int32(0)
	v355 = v217
	v360 = v216
	v362 = v5
	v367 = v216
	v368 = v217
	goto L22
L24:
	;
	goto L25
L25:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v226 = v19 + v220<<(uint(int32(2))%32) + int32(base.Ui32(v23)>>(uint(int32(12))%32))
	v232 = int32(base.Ui32(v23)>>(uint(int32(1))%32)) & int32(2047)
	if v232 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v233 = int32(-1)
	goto L28
L27:
	;
	v233 = int32(0)
	goto L28
L28:
	;
	__phi242 = v26
	__phi243 = int32(0)
	__phi244 = int32(1)
	v242 = __phi242
	v243 = __phi243
	v244 = __phi244
	goto L29
L29:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	if v252 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v355 = v343
	v360 = v347
	v362 = v348
	v367 = v349
	v368 = int32(base.Ui32(v344) >> (uint(int32(31)) % 32))
	goto L22
L31:
	;
	goto L30
L32:
	;
	v333 = int32(1)
	v335 = v244 + v333
	if v332 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L33:
	;
	if v252 < v232 {
		goto L64
	} else {
		goto L65
	}
L34:
	;
	v326 = int32(0)
	v343 = v243
	v344 = v326
	v347 = v244
	v348 = v242
	v349 = v326
	goto L31
L35:
	;
	v259 = v242 + int32(20)
	if base.Ui32(v252) < base.Ui32(v232) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	if v232 != 0 {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v255 = v233
	goto L38
L38:
	;
	if v255 == int32(0) {
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v255 = base.B2i32(int32(0) < v252)
	goto L38
L40:
	;
	v332 = v255
	goto L32
L41:
	;
	v261 = v252
	goto L43
L42:
	;
	v261 = v232
	goto L43
L43:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v261) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v323 != 0 {
		v332 = v323
		goto L32
	} else {
		goto L62
	}
L45:
	;
	v323 = int32(0)
	goto L44
L46:
	;
	v297 = v292
	v298 = v293
	v299 = v294
	goto L56
L47:
	;
	if (v259|v226)&int32(3) != 0 {
		v292 = v259
		v293 = v226
		v294 = v261
		goto L46
	} else {
		goto L50
	}
L48:
	;
	v285 = v259
	v286 = v226
	v287 = v261
	goto L49
L49:
	;
	if v287 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L50:
	;
	v269 = v259
	v270 = v226
	v271 = v261
	goto L51
L51:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v274 != v275 {
		v292 = v269
		v293 = v270
		v294 = v271
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v285 = v280
	v286 = v278
	v287 = v282
	goto L49
L53:
	;
	v277 = int32(4)
	v278 = v270 + v277
	v280 = v269 + v277
	v282 = v271 - v277
	if base.Ui32(int32(3)) < base.Ui32(v282) {
		v269 = v280
		v270 = v278
		v271 = v282
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v292 = v285
	v293 = v286
	v294 = v287
	goto L46
L56:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v302 == v303 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v323 = v302 - v303
	goto L44
L58:
	;
	v305 = int32(1)
	v310 = v299 - v305
	if v310 != 0 {
		v297 = v297 + v305
		v298 = v298 + v305
		v299 = v310
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L45
L62:
	;
	if v252 != v232 {
		goto L33
	} else {
		goto L63
	}
L63:
	;
	goto L34
L64:
	;
	v331 = int32(-1)
	goto L66
L65:
	;
	v331 = int32(1)
	goto L66
L66:
	;
	v332 = v331
	goto L32
L67:
	;
	v340 = int32(8)
	goto L69
L68:
	;
	v340 = int32(12)
	goto L69
L69:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v242+v340)))
	if v342 != 0 {
		__phi242 = v342
		__phi243 = v242
		__phi244 = v335
		v242 = __phi242
		v243 = __phi243
		v244 = __phi244
		goto L29
	} else {
		goto L70
	}
L70:
	;
	v343 = v242
	v344 = v332
	v347 = v335
	v348 = v5
	v349 = v333
	goto L31
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v360
	goto L73
L72:
	;
	goto L73
L73:
	;
	if v367 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v379 = F_MemoryContextAlloc(m, l0, int32(base.Ui32(v372)>>(uint(int32(1))%32))&int32(2047)+int32(20))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v409 + int32(1)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v413 + v203
	goto L1
L77:
	;
	return
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v379)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v379)+4)) = v203
	v384 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v390 = int32(base.Ui32(v386)>>(uint(v384)%32)) & int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v379)+16)) = v390
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v390 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v355 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v402 = F__emscripten_memcpy_bulkmem(m, v379+int32(20), v19+v394<<(uint(int32(2))%32)+int32(base.Ui32(v398)>>(uint(int32(12))%32)), v390)
	mBase = m.M
	goto L82
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v379
	return
L84:
	;
	goto L85
L85:
	;
	if v368 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v379
	return
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+12)) = v379
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 == int32(2249) {
		F_errstart_cold(m, int32(21), int32(581926))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v20
				F_errmsg(m, int32(112315), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errdetail(m, int32(658079), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errfinish(m, int32(523902), int32(7819), int32(414375))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
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
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v43 = F_make_expanded_record_from_typeid(m, v40, int32(-1), v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v43
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
				F_errmsg(m, int32(420724), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(1040), int32(335027))
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
				F_errmsg(m, int32(420724), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(1054), int32(313763))
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
				F_errmsg(m, int32(420724), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(1026), int32(315176))
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
				F_errmsg(m, int32(420400), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524456), int32(1150), int32(313716))
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
	v9 = v8 + v5
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
				F_errmsg(m, int32(420400), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524456), int32(1122), int32(315134))
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
				F_errmsg(m, int32(420378), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(958), int32(335035))
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(65535)
	v7 = v5 & v6
	if v7 != v6 {
		if v7 == int32(0) {
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
					F_errmsg(m, int32(251260), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515515), int32(1196), int32(442089))
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
			v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
			v14 = base.I32_rem_s(v12, base.I32_extend16_s(v5))
			v15 = v14
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
				F_errmsg(m, int32(420378), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(972), int32(313772))
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
	if v2&int32(65535) == int32(32768) {
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
				F_errmsg(m, int32(420378), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(922), int32(300840))
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
				F_errmsg(m, int32(420724), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(1115), int32(313789))
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
				F_errmsg(m, int32(420400), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524456), int32(1008), int32(313696))
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
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
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
					F_errmsg(m, int32(420724), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515515), int32(1292), int32(326942))
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
				__phi30 = v17
				v28 = __phi28
				v30 = __phi30
				for {
					v32 = base.I32_rem_s(v30, v28)
					if v32 != 0 {
						__phi28 = v32
						__phi30 = v28
						v28 = __phi28
						v30 = __phi30
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
			__phi30 = v17
			v28 = __phi28
			v30 = __phi30
			for {
				v32 = base.I32_rem_s(v30, v28)
				if v32 != 0 {
					__phi28 = v32
					__phi30 = v28
					v28 = __phi28
					v30 = __phi30
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
				F_errmsg(m, int32(420724), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(909), int32(511908))
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
				F_errmsg(m, int32(420724), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515515), int32(829), int32(315142))
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v9) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v9 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v12>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
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
				F_errmsg(m, int32(420400), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524456), int32(1069), int32(313780))
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
					F_errmsg(m, int32(420400), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524456), int32(774), int32(513962))
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
					F_errmsg(m, int32(420400), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524456), int32(787), int32(513962))
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
					F_errmsg(m, int32(420400), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524456), int32(736), int32(511900))
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
					F_errmsg(m, int32(420400), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524456), int32(750), int32(511900))
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
							F_errmsg(m, int32(420400), int32(0))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524456), int32(704), int32(305286))
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
								F_errmsg(m, int32(420400), int32(0))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(524456), int32(710), int32(305286))
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
							F_errmsg(m, int32(420400), int32(0))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524456), int32(636), int32(326925))
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
								F_errmsg(m, int32(420400), int32(0))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(524456), int32(704), int32(305286))
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
									F_errmsg(m, int32(420400), int32(0))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(524456), int32(710), int32(305286))
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = v9 - v5
	if base.B2i32(int64(0) < v5) != base.B2i32(v10 < v9) {
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
				F_errmsg(m, int32(420400), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524456), int32(485), int32(334951))
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
		v31 = F_Int64GetDatum(m, v10)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
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
					F_errmsg(m, int32(251260), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524456), int32(572), int32(442073))
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = v9 + v5
	if base.B2i32(v5 < int64(0)) != base.B2i32(v10 < v9) {
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
				F_errmsg(m, int32(420400), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524456), int32(471), int32(315100))
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
		v31 = F_Int64GetDatum(m, v10)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
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
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
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
	v14 = F_ArrayGetNItems(m, v11, l0+int32(16))
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
	F_errmsg(m, int32(161982), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(520866), int32(344), int32(73817))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
	var v66 int32
	_ = v66
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v315 int32
	_ = v315
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v801 int32
	_ = v801
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v916 int32
	_ = v916
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
		v740 = v14
		v742 = v14
		v747 = v14
		v749 = v14
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v757 = F_buildoidvector(m, v50, v747)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L201
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v62 <= int32(0) {
		v740 = v14
		v742 = v14
		v747 = v14
		v749 = v14
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v66 = base.B2i32(l3 != int32(29))
	v88 = v14
	v89 = v14
	v90 = v14
	v91 = v14
	v96 = v14
	v98 = v14
	goto L22
L12:
	;
	v740 = base.B2i32(int32(0) < v285)
	v742 = v269
	v747 = v232
	v749 = v461
	goto L9
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L196
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L4
	} else {
		goto L191
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L186
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L181
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L4
	} else {
		goto L176
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L171
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L4
	} else {
		goto L166
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L160
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L154
	}
L22:
	;
	v107 = v89 << (uint(int32(2)) % 32)
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
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L148
	}
L24:
	;
	if v114 == int32(0) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+22)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v119)+82)))
	if v121 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L23
L27:
	;
	v151 = F_typeTypeId(m, v114)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L38
	}
L28:
	;
	if base.B2i32(l2 != int32(14)) == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	if l3 == int32(1) {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v128 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v128 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v135 = F_TypeNameToString(m, v112)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v135
	F_errmsg(m, int32(318228), v40-int32(-64))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(517209), int32(259), int32(79769))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	F_ReleaseCatCache(m, v114)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v159 = F_object_aclcheck(m, int32(1247), v151, v157, int64(256))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v159 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_aclcheck_error_type(m, v159, v151)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+12)))
	if v163 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v215 = v111 - int32(111)
	switch v215 {
	case 0, 5:
		v231 = int32(0)
		v232 = v96
		goto L64
	default:
		goto L65
	}
L48:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if l3 != int32(29) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_errmsg(m, int32(128873), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L61
	}
L51:
	;
	if l3 != int32(1) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_errmsg(m, int32(128950), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L58
	}
L54:
	;
	F_errmsg(m, int32(128911), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(517209), int32(284), int32(79769))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(517209), int32(289), int32(79769))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(517209), int32(294), int32(79769))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	if v111 == int32(100) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	if int32(0) < v90 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v96<<(uint(int32(2))%32)))) = v151
	v222 = int32(1)
	v224 = v96 + v222
	if l5 == int32(0) {
		v231 = v222
		v232 = v224
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v228 = F_lappend_oid(m, v227, v151)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v228
	v231 = v222
	v232 = v224
	goto L64
L69:
	;
	if v236 != int32(118) {
		v285 = v90
		goto L86
	} else {
		goto L87
	}
L70:
	;
	if v66 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v236 = int32(105)
	goto L73
L72:
	;
	v236 = v111
	goto L73
L73:
	;
	v238 = v236 - int32(105)
	switch v238 {
	case 0, 13:
		v269 = v91
		goto L69
	default:
		goto L70
	}
L74:
	;
	v269 = v91 + int32(1)
	goto L69
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v263
	goto L74
L76:
	;
	if v90 <= int32(0) {
		v263 = int32(2249)
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v91 != 0 {
		goto L74
	} else {
		goto L85
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(227294), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(517209), int32(326), int32(79769))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v263 = v151
	goto L75
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107+v52))) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v107+v54))) = base.I32_extend8_s(v236)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v291 == int32(0) {
		v461 = v98
		goto L92
	} else {
		goto L93
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v151
	v274 = v90 + int32(1)
	if base.Ui32(v151-int32(2276)) < base.Ui32(int32(2)) {
		v285 = v274
		goto L86
	} else {
		goto L88
	}
L88:
	;
	if v151 == int32(5078) {
		v285 = v274
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v281 = F_get_element_type(m, v151)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	if v281 == int32(0) {
		goto L18
	} else {
		goto L91
	}
L91:
	;
	v285 = v274
	goto L86
L92:
	;
	if l9 != 0 {
		goto L121
	} else {
		goto L122
	}
L93:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v294 == int32(0) {
		v461 = v98
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v297 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v428 = F_cstring_to_text(m, v291)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L120
	}
L96:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v315 = int32(0)
	goto L97
L97:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v300+v315<<(uint(int32(2))%32))))
	if v342 == v110 {
		goto L95
	} else {
		goto L99
	}
L98:
	;
	goto L95
L99:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	if v345 == int32(100) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v348 = int32(105)
	goto L102
L101:
	;
	v348 = v345
	goto L102
L102:
	;
	switch v238 {
	case 0, 13:
		goto L107
	default:
		goto L106
	}
L103:
	;
	v388 = v315 + int32(1)
	if v388 != v297 {
		v315 = v388
		goto L97
	} else {
		goto L119
	}
L104:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	if v353 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L105:
	;
	switch v215 {
	case 0, 5:
		goto L103
	default:
		goto L104
	}
L106:
	;
	switch v348 - int32(105) {
	case 0, 13:
		goto L105
	default:
		goto L104
	}
L107:
	;
	switch v348 - int32(105) {
	case 0, 13:
		goto L105
	default:
		goto L104
	case 6, 11:
		goto L103
	}
L108:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v356 == int32(0) {
		goto L103
	} else {
		goto L109
	}
L109:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v362 == int32(0) {
		v381 = v361
		v382 = v362
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v382-v381 == int32(0) {
		goto L17
	} else {
		goto L118
	}
L111:
	;
	goto L110
L112:
	;
	if v361 != v362 {
		v381 = v361
		v382 = v362
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v366 = v353
	v367 = v291
	goto L114
L114:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)))
	if v371 == int32(0) {
		v381 = v370
		v382 = v371
		goto L111
	} else {
		goto L116
	}
L115:
	;
	v381 = v370
	v382 = v371
	goto L111
L116:
	;
	v374 = int32(1)
	if v370 == v371 {
		v366 = v366 + v374
		v367 = v367 + v374
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L103
L119:
	;
	goto L98
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+v107))) = v428
	v461 = int32(1)
	goto L92
L121:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v470 != 0 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if v481 != 0 {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	v474 = v470
	goto L126
L125:
	;
	v472 = F_pstrdup(m, int32(790230))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L127
	}
L126:
	;
	v475 = F_makeString(m, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L128
	}
L127:
	;
	v474 = v472
	goto L126
L128:
	;
	v477 = F_lappend(m, v469, v475)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v477
	goto L123
L130:
	;
	v512 = v89 + int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v513 <= v512 {
		goto L12
	} else {
		goto L147
	}
L131:
	;
	if v231 == int32(0) {
		goto L16
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	if v88&v231 != 0 {
		goto L14
	} else {
		goto L142
	}
L134:
	;
	v485 = F_transformExpr(m, l0, v481, int32(31))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v488 = F_coerce_to_specific_type(m, l0, v485, v151, int32(545230))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_assign_expr_collations(m, l0, v488)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v492 != 0 {
		goto L15
	} else {
		goto L138
	}
L138:
	;
	v493 = F_contain_var_clause(m, v488)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	if v493 != 0 {
		goto L15
	} else {
		goto L140
	}
L140:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v496 = F_lappend(m, v495, v488)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v496
	v510 = int32(1)
	goto L130
L142:
	;
	v501 = v88 & v66
	if l3 != int32(29) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v510 = v501
	goto L130
L144:
	;
	goto L145
L145:
	;
	v502 = int32(1)
	if (v88^v502)&v502 == int32(0) {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	v510 = v501
	goto L130
L147:
	;
	v88 = v510
	v89 = v512
	v90 = v285
	v91 = v269
	v96 = v232
	v98 = v461
	goto L22
L148:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v522 = F_TypeNameToString(m, v112)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v522
	F_errmsg(m, int32(202009), v40+int32(32))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(517209), int32(246), int32(79769))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v545 = F_TypeNameToString(m, v112)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v545
	F_errmsg(m, int32(202050), v40+int32(48))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(517209), int32(253), int32(79769))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v568 = F_TypeNameToString(m, v112)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v568
	F_errmsg(m, int32(76026), v40)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	F_parser_errposition(m, l0, v574)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(517209), int32(270), int32(79769))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
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
	F_errcode(m, int32(50724996))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	F_errmsg(m, int32(227242), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(517209), int32(305), int32(79769))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
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
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(26021), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(517209), int32(352), int32(79769))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
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
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v627
	F_errmsg(m, int32(434279), v40+int32(16))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(517209), int32(399), int32(79769))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
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
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	F_errmsg(m, int32(167411), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(517209), int32(417), int32(79769))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
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
	F_errcode(m, int32(393348))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errmsg(m, int32(362222), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(517209), int32(433), int32(79769))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
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
	F_errcode(m, int32(50724996))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errmsg(m, int32(131365), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(517209), int32(458), int32(79769))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
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
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(362394), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_parser_errposition(m, l0, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(517209), int32(469), int32(79769))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v757
	v760 = int32(0)
	if base.B2i32(v740 == v760)&base.B2i32(v742 <= v760) == v760 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v749&int32(1) != 0 {
		goto L209
	} else {
		goto L210
	}
L203:
	;
	v768 = F_construct_array_builtin(m, v52, v43, int32(26))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v779 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v779
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v779
	goto L202
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v768
	v772 = F_construct_array_builtin(m, v54, v43, int32(18))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v772
	if v742 < int32(2) {
		goto L202
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = int32(2249)
	goto L202
L209:
	;
	if int32(0) < v43 {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	v916 = int32(0)
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v916
	m.G0 = v40 + int32(80)
	return
L212:
	;
	v801 = int32(0)
	goto L215
L213:
	;
	goto L214
L214:
	;
	v876 = F_construct_array_builtin(m, v56, v43, int32(25))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L4
	} else {
		goto L222
	}
L215:
	;
	v827 = v56 + v801<<(uint(int32(2))%32)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)))
	if v828 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L214
L217:
	;
	v832 = F_cstring_to_text(m, int32(790230))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v836 = v801 + int32(1)
	if v836 != v43 {
		v801 = v836
		goto L215
	} else {
		goto L221
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v827))) = v832
	goto L219
L221:
	;
	goto L216
L222:
	;
	v916 = v876
	goto L211
}
func F_inv_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
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
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(2224)
	m.G0 = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v25&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L20
	} else {
		goto L111
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L20
	} else {
		goto L108
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L20
	} else {
		goto L104
	}
L4:
	;
	if int32(0) < l2 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L20
	} else {
		goto L100
	}
L7:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(4398046509057)) <= base.Ui64(v30+base.I64_extend_i32_u(l2)) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v364 = v4
	goto L9
L9:
	;
	m.G0 = v23 + int32(2224)
	return v364
L10:
	;
	v37 = base.I32_wrap_i64(int64(base.Ui64(v30) >> (uint(int64(11)) % 64)))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v42 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v39
	goto L13
L12:
	;
	v43 = int32(0)
	goto L13
L13:
	;
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = int32(4554292)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v50
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v77 = v39
	goto L16
L16:
	;
	v80 = v23 + int32(80)
	v81 = F_CatalogOpenIndexes(m, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L26
	}
L17:
	;
	v62 = v39
	v63 = v42
	goto L19
L18:
	;
	v55 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v63 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[802])) = v55
	v61 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	v62 = v55
	v63 = v61
	goto L19
L22:
	;
	v69 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v74 = v62
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v47
	v77 = v74
	goto L16
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[803])) = v69
	v73 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v74 = v73
	goto L24
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v23+int32(2128), int32(1), int32(3), int32(184), v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	F_ScanKeyInit(m, v23+int32(2176), int32(2), int32(4), int32(150), v37)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v99 = v23 + int32(80)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v107 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v112 = F_systable_beginscan_ordered(m, v105, v107, v108, int32(2), v23+int32(2128))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v122 = int32(1)
	v123 = v4
	v125 = int32(0)
	v126 = v37
	v128 = v4
	goto L30
L30:
	;
	if v122&int32(1) == int32(0) {
		v153 = v123
		v154 = v128
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_systable_endscan_ordered(m, v112)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L20
	} else {
		goto L97
	}
L32:
	;
	if v153 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v141 = F_systable_getnext_ordered(m, v112, int32(1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	if v141 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v153 = v123
	v154 = int32(0)
	goto L32
L36:
	;
	goto L37
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+20)))
	if v147&int32(1) != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v153 = v146 + v150
	v154 = v141
	goto L32
L39:
	;
	F_pfree(m, v334)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L20
	} else {
		goto L95
	}
L40:
	;
	v270 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v271 = base.I32_wrap_i64(v270)
	v273 = v271 & int32(2047)
	if v273 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v157 != v126 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v160 = v153 + int32(8)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+8)))
	v163 = v161 & int32(3)
	if v163 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = F_detoast_attr(m, v160)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L20
	} else {
		goto L46
	}
L44:
	;
	v166 = v160
	goto L45
L45:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v169 = int32(base.Ui32(v167) >> (uint(int32(2)) % 32))
	v171 = v169 - int32(4)
	if base.Ui32(v167-int32(8212)) <= base.Ui32(int32(-8197)) {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v166 = v164
	goto L45
L47:
	;
	if v171 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v163 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v178 = F__emscripten_memcpy_bulkmem(m, v80, v166+int32(4), v171)
	mBase = m.M
	v179 = v178
	goto L51
L50:
	;
	v179 = v80
	goto L51
L51:
	;
	goto L48
L52:
	;
	F_pfree(m, v166)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v185 = base.I32_wrap_i64(v182) & int32(2047)
	if v185 <= v171 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v220 = int32(2048) - v185
	v221 = l2 - v125
	if v220 < v221 {
		goto L67
	} else {
		goto L68
	}
L57:
	;
	v187 = v185 - v171
	v190 = v23 + int32(76) + v169
	if v190&int32(3) != 0 {
		v209 = v187
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v213 = F__emscripten_memset_bulkmem(m, v190, base.I32_extend8_s(int32(0)), v209)
	mBase = m.M
	goto L66
L59:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v187) {
		v209 = v187
		goto L58
	} else {
		goto L60
	}
L60:
	;
	if v187&int32(3) != 0 {
		v209 = v187
		goto L58
	} else {
		goto L61
	}
L61:
	;
	if base.Ui32(v179+v185) <= base.Ui32(v190) {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v199 = v169 + v99
	v200 = v185 + v99
	if base.Ui32(v200) < base.Ui32(v199) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v202 = v199
	goto L65
L64:
	;
	v202 = v200
	goto L65
L65:
	;
	v209 = (v202+(v23+int32(76)^int32(-1))-v169)&int32(-4) + int32(4)
	goto L58
L66:
	;
	goto L56
L67:
	;
	v223 = v220
	goto L69
L68:
	;
	v223 = v221
	goto L69
L69:
	;
	if v223 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v182 + base.I64_extend_i32_s(v223)
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+62)) = uint8(v229)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+60)) = uint16(v229)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+56)) = uint16(v229)
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+58)) = uint8(v238)
	v241 = v223 + v185
	if v241 < v171 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v224 = F__emscripten_memcpy_bulkmem(m, v179+v185, l1+v125, v223)
	mBase = m.M
	goto L73
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v243 = v171
	goto L76
L75:
	;
	v243 = v241
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v243<<(uint(int32(2))%32) + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v23 + int32(76)
	v253 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+52))
	v261 = F_heap_modify_tuple(m, v154, v254, v23-int32(-64), v23+int32(60), v23+int32(56))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L20
	} else {
		goto L77
	}
L77:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	F_CatalogTupleUpdateWithInfo(m, v264, v261+int32(4), v261, v81)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	v334 = v261
	v335 = v223
	v336 = v238
	v337 = v229
	v339 = int32(0)
	goto L39
L79:
	;
	v295 = int32(2048) - v273
	v296 = l2 - v125
	if v295 < v296 {
		goto L86
	} else {
		goto L87
	}
L80:
	;
	if v271&int32(3) != 0 {
		v289 = v273
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v291 = F__emscripten_memset_bulkmem(m, v80, base.I32_extend8_s(int32(0)), v289)
	mBase = m.M
	goto L85
L82:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v273) {
		v289 = v273
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if base.Ui32(v273+v80) <= base.Ui32(v80) {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v289 = (v273-int32(1))&int32(-4) + int32(4)
	goto L81
L85:
	;
	goto L79
L86:
	;
	v298 = v295
	goto L88
L87:
	;
	v298 = v296
	goto L88
L88:
	;
	if v298 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v270 + base.I64_extend_i32_s(v298)
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+62)) = uint8(v304)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+60)) = uint16(v304)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = (v273+v298)<<(uint(int32(2))%32) + int32(16)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v23 + int32(76)
	v322 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+52))
	v328 = F_heap_form_tuple(m, v323, v23-int32(-64), v23+int32(60))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L20
	} else {
		goto L93
	}
L90:
	;
	v299 = F__emscripten_memcpy_bulkmem(m, v273+v80, l1+v125, v298)
	mBase = m.M
	goto L92
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	F_CatalogTupleInsertWithInfo(m, v331, v328, v81)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L20
	} else {
		goto L94
	}
L94:
	;
	v334 = v328
	v335 = v298
	v336 = v304
	v337 = v153
	v339 = v154
	goto L39
L95:
	;
	v347 = v335 + v125
	if v347 < l2 {
		v122 = v336
		v123 = v337
		v125 = v347
		v126 = v126 + int32(1)
		v128 = v339
		goto L30
	} else {
		goto L96
	}
L96:
	;
	goto L31
L97:
	;
	F_CatalogCloseIndexes(m, v81)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L20
	} else {
		goto L99
	}
L99:
	;
	v364 = v347
	goto L9
L100:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L20
	} else {
		goto L101
	}
L101:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v386
	F_errmsg(m, int32(45664), v23)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L20
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(521291), int32(580), int32(366444))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L20
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l2
	F_errmsg(m, int32(505240), v23+int32(16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L20
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(521291), int32(590), int32(366444))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L20
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errmsg_internal(m, int32(117672), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L20
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(521291), int32(624), int32(366444))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L20
	} else {
		goto L112
	}
L112:
	;
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v434
	F_errmsg(m, int32(496360), v23+int32(32))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L20
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(521291), int32(153), int32(451420))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1258]))
	if base.Ui32(l0) < base.Ui32(v10) {
		v27 = v2
		m.G0 = v6 + int32(16)
		return v27
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
		if base.Ui32(v13) < base.Ui32(l0) {
			v27 = v2
			m.G0 = v6 + int32(16)
			return v27
		} else {
			v21 = F_bsearch(m, v6+int32(12), int32(1900272), int32(34), int32(8), int32(4663))
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
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4 == v2 {
		v24 = v2
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
			v24 = v2
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
			if v19 != 0 {
				v24 = v2
			} else {
				v24 = int32(1)
			}
		}
	}
	return v24
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_string2ean(m, v7, v8, v5+int32(8), int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v24 = int32(0)
			m.G0 = v5 + int32(16)
			return v24
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			v22 = F_Int64GetDatum(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				m.G0 = v5 + int32(16)
				return v24
			}
		}
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	F_ean2isn(m, v8, v5+int32(8), int32(5))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v17 = F_Int64GetDatum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v17
		}
	}
}
func F_iswalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	if base.Ui32(l0) <= base.Ui32(int32(131071)) {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_consts[641]))))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v12<<(uint(int32(5))%32))+uint32(_consts[641]))))
		return int32(base.Ui32(v18)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		return base.B2i32(base.Ui32(l0) < base.Ui32(int32(196606)))
	}
}
func F_iswgraph(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v2 = int32(0)
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	v46 = int32(0)
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
	v46 = base.B2i32(v39 != int32(0))
	goto L1
L6:
	;
	v12 = int32(4130208)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v22 = int32(4130208)
	goto L18
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v14 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v14 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if l0 != v14 {
		v12 = v12 + int32(4)
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
	v20 = v12
	goto L17
L16:
	;
	v20 = int32(0)
	goto L17
L17:
	;
	v39 = v20
	goto L5
L18:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v28 != 0 {
		v22 = v22 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v29 = int32(4130208)
	v39 = (v22-v29)&int32(-4) + v29
	goto L5
L20:
	;
	goto L19
L21:
	;
	v81 = v2
	goto L23
L22:
	;
	if base.Ui32(l0) <= base.Ui32(int32(254)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	return v81
L24:
	;
	v81 = base.B2i32(v78 != int32(0))
	goto L23
L25:
	;
	v78 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	goto L24
L26:
	;
	goto L27
L27:
	;
	v56 = int32(1)
	if base.Ui32(l0-int32(57344)) < base.Ui32(int32(8185)) {
		v76 = v56
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v78 = v76
	goto L24
L29:
	;
	if base.Ui32(l0) < base.Ui32(int32(8232)) {
		v76 = v56
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(l0-int32(8234)) < base.Ui32(int32(47062)) {
		v76 = v56
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v67 = int32(65534)
	v76 = base.B2i32(l0&v67 != v67) & base.B2i32(base.Ui32(l0-int32(65532)) < base.Ui32(int32(1048580)))
	goto L28
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
	var v154 int32
	_ = v154
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
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
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
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2536 int32
	_ = v2536
	var v2540 int32
	_ = v2540
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2565 int32
	_ = v2565
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2594 int32
	_ = v2594
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2611 int32
	_ = v2611
	var v2627 int32
	_ = v2627
	var v2631 int32
	_ = v2631
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2700 int32
	_ = v2700
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2713 int32
	_ = v2713
	var v2719 int32
	_ = v2719
	var v2735 int32
	_ = v2735
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2764 int32
	_ = v2764
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2790 int32
	_ = v2790
	var v2794 int32
	_ = v2794
	var v2810 int32
	_ = v2810
	var v2814 int32
	_ = v2814
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2834 int32
	_ = v2834
	var v2836 int32
	_ = v2836
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2861 int32
	_ = v2861
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2883 int32
	_ = v2883
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2896 int32
	_ = v2896
	var v2902 int32
	_ = v2902
	var v2918 int32
	_ = v2918
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2969 int32
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2983 int32
	_ = v2983
	var v2985 int32
	_ = v2985
	var v2990 int32
	_ = v2990
	var v2992 int32
	_ = v2992
	var v2999 int32
	_ = v2999
	var v3002 int32
	_ = v3002
	var v3006 int32
	_ = v3006
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3028 int32
	_ = v3028
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3039 int32
	_ = v3039
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
	v18 = F_memcmp(m, v16+v6, int32(2242088), v8)
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
	v32 = F_slice_from_s(m, l0, int32(5), int32(2242094))
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
	return v3039
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
	v50 = F_find_among(m, l0, int32(4340256), int32(7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v154 = v6
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
	v88 = F_slice_from_s(m, l0, int32(2), int32(2242109))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L36
	}
L21:
	;
	v82 = F_slice_from_s(m, l0, int32(2), int32(2242107))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L34
	}
L22:
	;
	v76 = F_slice_from_s(m, l0, int32(2), int32(2242105))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L32
	}
L23:
	;
	v70 = F_slice_from_s(m, l0, int32(2), int32(2242103))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L30
	}
L24:
	;
	v64 = F_slice_from_s(m, l0, int32(2), int32(2242101))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L28
	}
L25:
	;
	v58 = F_slice_from_s(m, l0, int32(2), int32(2242099))
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
	v3039 = v58
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
	v3039 = v64
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
	v3039 = v70
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
	v3039 = v76
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
	v3039 = v82
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
	v3039 = v88
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
	v2611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2611
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2611
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L655
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L67
L63:
	;
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2589
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2587)+8))
	if v2589 < v2591 {
		goto L61
	} else {
		goto L644
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
	if v169 <= v154 {
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
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v170))))
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
	v191 = v154 + int32(1)
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
	v200 = v154 + int32(2)
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
	v244 = v229&int32(63) | (v187<<(uint(int32(18))%32)&int32(1835008) | v196<<(uint(int32(12))%32) | v212<<(uint(int32(6))%32))
	v245 = int32(4)
	goto L72
L81:
	;
	v216 = v154 + int32(3)
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
	v244 = v187<<(uint(int32(12))%32)&int32(61440) | v196<<(uint(int32(6))%32) | v212
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
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v249)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
	if int32(base.Ui32(v255)>>(uint(v249&int32(7))%32))&int32(1) == int32(0) {
		v267 = v245
		goto L66
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v245 + v154
	goto L88
L88:
	;
	goto L68
L89:
	;
	v2583 = F_slice_from_s(m, l0, int32(1), int32(2242133))
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		goto L8
	} else {
		goto L642
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
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
	v375 = v360&int32(63) | (v318<<(uint(int32(18))%32)&int32(1835008) | v327<<(uint(int32(12))%32) | v343<<(uint(int32(6))%32))
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
	v375 = v318<<(uint(int32(12))%32)&int32(61440) | v327<<(uint(int32(6))%32) | v343
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
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v380)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v410 = F_slice_from_s(m, l0, int32(1), int32(2242132))
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
		v3039 = v410
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
	v514 = v499&int32(63) | (v457<<(uint(int32(18))%32)&int32(1835008) | v466<<(uint(int32(12))%32) | v482<<(uint(int32(6))%32))
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
	v514 = v457<<(uint(int32(12))%32)&int32(61440) | v466<<(uint(int32(6))%32) | v482
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
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v519)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
		v154 = v603
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
	v558 = v154
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
	*(*int32)(unsafe.Add(mBase, uint32(v1753)+8)) = v1751
	goto L174
L176:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1751 = v1749 + v1747
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
	v701 = v686&int32(63) | (v644<<(uint(int32(18))%32)&int32(1835008) | v653<<(uint(int32(12))%32) | v669<<(uint(int32(6))%32))
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
	v701 = v644<<(uint(int32(12))%32)&int32(61440) | v653<<(uint(int32(6))%32) | v669
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
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v706)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v820 = v805&int32(63) | (v763<<(uint(int32(18))%32)&int32(1835008) | v772<<(uint(int32(12))%32) | v788<<(uint(int32(6))%32))
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
	v820 = v763<<(uint(int32(12))%32)&int32(61440) | v772<<(uint(int32(6))%32) | v788
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
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v825)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
		v1747 = v968
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
	v939 = v924&int32(63) | (v882<<(uint(int32(18))%32)&int32(1835008) | v891<<(uint(int32(12))%32) | v907<<(uint(int32(6))%32))
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
	v939 = v882<<(uint(int32(12))%32)&int32(61440) | v891<<(uint(int32(6))%32) | v907
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
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v944)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v1060 = v1045&int32(63) | (v1003<<(uint(int32(18))%32)&int32(1835008) | v1012<<(uint(int32(12))%32) | v1028<<(uint(int32(6))%32))
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
	v1060 = v1003<<(uint(int32(12))%32)&int32(61440) | v1012<<(uint(int32(6))%32) | v1028
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
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1065)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
		v1747 = v1208
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
	v1178 = v1163&int32(63) | (v1121<<(uint(int32(18))%32)&int32(1835008) | v1130<<(uint(int32(12))%32) | v1146<<(uint(int32(6))%32))
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
	v1178 = v1121<<(uint(int32(12))%32)&int32(61440) | v1130<<(uint(int32(6))%32) | v1146
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
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1183)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v1301 = v1286&int32(63) | (v1244<<(uint(int32(18))%32)&int32(1835008) | v1253<<(uint(int32(12))%32) | v1269<<(uint(int32(6))%32))
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
	v1301 = v1244<<(uint(int32(12))%32)&int32(61440) | v1253<<(uint(int32(6))%32) | v1269
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
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1306)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v1419 = v1404&int32(63) | (v1362<<(uint(int32(18))%32)&int32(1835008) | v1371<<(uint(int32(12))%32) | v1387<<(uint(int32(6))%32))
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
	v1419 = v1362<<(uint(int32(12))%32)&int32(61440) | v1371<<(uint(int32(6))%32) | v1387
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
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1424)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
		v1747 = v1567
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
	v1538 = v1523&int32(63) | (v1481<<(uint(int32(18))%32)&int32(1835008) | v1490<<(uint(int32(12))%32) | v1506<<(uint(int32(6))%32))
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
	v1538 = v1481<<(uint(int32(12))%32)&int32(61440) | v1490<<(uint(int32(6))%32) | v1506
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
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1543)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v1659 = v1644&int32(63) | (v1602<<(uint(int32(18))%32)&int32(1835008) | v1611<<(uint(int32(12))%32) | v1627<<(uint(int32(6))%32))
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
	v1659 = v1602<<(uint(int32(12))%32)&int32(61440) | v1611<<(uint(int32(6))%32) | v1627
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
	v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1664)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
		v1751 = v1744
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
	v1845 = v1830&int32(63) | (v1788<<(uint(int32(18))%32)&int32(1835008) | v1797<<(uint(int32(12))%32) | v1813<<(uint(int32(6))%32))
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
	v1845 = v1788<<(uint(int32(12))%32)&int32(61440) | v1797<<(uint(int32(6))%32) | v1813
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
	v1856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1850)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v1967 = v1952&int32(63) | (v1910<<(uint(int32(18))%32)&int32(1835008) | v1919<<(uint(int32(12))%32) | v1935<<(uint(int32(6))%32))
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
	v1967 = v1910<<(uint(int32(12))%32)&int32(61440) | v1919<<(uint(int32(6))%32) | v1935
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
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1972)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v2092 = v2077&int32(63) | (v2035<<(uint(int32(18))%32)&int32(1835008) | v2044<<(uint(int32(12))%32) | v2060<<(uint(int32(6))%32))
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
	v2092 = v2035<<(uint(int32(12))%32)&int32(61440) | v2044<<(uint(int32(6))%32) | v2060
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
	v2103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2097)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v2214 = v2199&int32(63) | (v2157<<(uint(int32(18))%32)&int32(1835008) | v2166<<(uint(int32(12))%32) | v2182<<(uint(int32(6))%32))
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
	v2214 = v2157<<(uint(int32(12))%32)&int32(61440) | v2166<<(uint(int32(6))%32) | v2182
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
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2219)>>(uint(int32(3))%32)))+uint32(_consts[1308]))))
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
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2312
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2312
	v2317 = F_find_among_b(m, l0, int32(4341264), int32(51))
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L8
	} else {
		goto L553
	}
L537:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259+v2257))))
	if v2261&int32(224) != int32(96) {
		goto L536
	} else {
		goto L538
	}
L538:
	;
	if int32(1)<<(uint(v2261)%32)&int32(33314) == int32(0) {
		goto L536
	} else {
		goto L539
	}
L539:
	;
	v2274 = F_find_among_b(m, l0, int32(4340400), int32(37))
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L8
	} else {
		goto L540
	}
L540:
	;
	if v2274 == int32(0) {
		goto L536
	} else {
		goto L541
	}
L541:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2278
	v2281 = v2278 - int32(1)
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2281 <= v2282 {
		goto L536
	} else {
		goto L542
	}
L542:
	;
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2284+v2281))))
	switch v2286 - int32(111) {
	case 0, 3:
		goto L543
	default:
		goto L536
	}
L543:
	;
	v2291 = F_find_among_b(m, l0, int32(4341152), int32(5))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L8
	} else {
		goto L544
	}
L544:
	;
	if v2291 == int32(0) {
		goto L536
	} else {
		goto L545
	}
L545:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+8))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2297 < v2296 {
		goto L536
	} else {
		goto L546
	}
L546:
	;
	switch v2291 - int32(1) {
	case 0:
		goto L548
	case 1:
		goto L547
	default:
		goto L536
	}
L547:
	;
	v2307 = F_slice_from_s(m, l0, int32(1), int32(2242146))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L8
	} else {
		goto L551
	}
L548:
	;
	v2301 = F_slice_del(m, l0)
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L8
	} else {
		goto L549
	}
L549:
	;
	if int32(0) <= v2301 {
		goto L536
	} else {
		goto L550
	}
L550:
	;
	v3039 = v2301
	goto L13
L551:
	;
	if v2307 < int32(0) {
		v3039 = v2307
		goto L13
	} else {
		goto L552
	}
L552:
	;
	goto L536
L553:
	;
	if v2317 == int32(0) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2587 = v2321
	goto L64
L555:
	;
	goto L556
L556:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2322
	switch v2317 - int32(1) {
	case 0:
		goto L565
	case 1:
		goto L564
	case 2:
		goto L563
	case 3:
		goto L562
	case 4:
		goto L561
	case 5:
		goto L560
	case 6:
		goto L559
	case 7:
		goto L558
	case 8:
		goto L557
	default:
		goto L61
	}
L557:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(v2516)))
	if v2322 < v2517 {
		v2587 = v2516
		goto L64
	} else {
		goto L623
	}
L558:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v2475)))
	if v2322 < v2476 {
		v2587 = v2475
		goto L64
	} else {
		goto L612
	}
L559:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v2403)+4))
	if v2322 < v2404 {
		v2587 = v2403
		goto L64
	} else {
		goto L592
	}
L560:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+8))
	if v2322 < v2397 {
		v2587 = v2396
		goto L64
	} else {
		goto L589
	}
L561:
	;
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2387)))
	if v2322 < v2388 {
		v2587 = v2387
		goto L64
	} else {
		goto L586
	}
L562:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2378)))
	if v2322 < v2379 {
		v2587 = v2378
		goto L64
	} else {
		goto L583
	}
L563:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v2369)))
	if v2322 < v2370 {
		v2587 = v2369
		goto L64
	} else {
		goto L580
	}
L564:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2333)))
	if v2322 < v2334 {
		v2587 = v2333
		goto L64
	} else {
		goto L569
	}
L565:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2326)))
	if v2322 < v2327 {
		v2587 = v2326
		goto L64
	} else {
		goto L566
	}
L566:
	;
	v2329 = F_slice_del(m, l0)
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L8
	} else {
		goto L567
	}
L567:
	;
	if int32(0) <= v2329 {
		goto L61
	} else {
		goto L568
	}
L568:
	;
	v3039 = v2329
	goto L13
L569:
	;
	v2336 = F_slice_del(m, l0)
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L8
	} else {
		goto L570
	}
L570:
	;
	if v2336 < int32(0) {
		v3039 = v2336
		goto L13
	} else {
		goto L571
	}
L571:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2340
	v2342 = int32(2)
	v2344 = int32(0)
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2340-v2347 < v2342 {
		v2357 = v2344
		goto L573
	} else {
		goto L574
	}
L572:
	;
	if v2357 == int32(0) {
		goto L61
	} else {
		goto L576
	}
L573:
	;
	goto L572
L574:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2353 = F_memcmp(m, v2350+v2340-v2342, int32(2242298), v2342)
	mBase = m.M
	if v2353 != 0 {
		v2357 = v2344
		goto L573
	} else {
		goto L575
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2340 - v2342
	v2357 = int32(1)
	goto L573
L576:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2360
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2362)))
	if v2360 < v2363 {
		goto L61
	} else {
		goto L577
	}
L577:
	;
	v2365 = F_slice_del(m, l0)
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L8
	} else {
		goto L578
	}
L578:
	;
	if int32(0) <= v2365 {
		goto L61
	} else {
		goto L579
	}
L579:
	;
	v3039 = v2365
	goto L13
L580:
	;
	v2374 = F_slice_from_s(m, l0, int32(3), int32(2242300))
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L8
	} else {
		goto L581
	}
L581:
	;
	if int32(0) <= v2374 {
		goto L61
	} else {
		goto L582
	}
L582:
	;
	v3039 = v2374
	goto L13
L583:
	;
	v2383 = F_slice_from_s(m, l0, int32(1), int32(2242303))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L8
	} else {
		goto L584
	}
L584:
	;
	if int32(0) <= v2383 {
		goto L61
	} else {
		goto L585
	}
L585:
	;
	v3039 = v2383
	goto L13
L586:
	;
	v2392 = F_slice_from_s(m, l0, int32(4), int32(2242304))
	mBase = m.M
	v2393 = m.ExcPending
	if v2393 != 0 {
		goto L8
	} else {
		goto L587
	}
L587:
	;
	if int32(0) <= v2392 {
		goto L61
	} else {
		goto L588
	}
L588:
	;
	v3039 = v2392
	goto L13
L589:
	;
	v2399 = F_slice_del(m, l0)
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L8
	} else {
		goto L590
	}
L590:
	;
	if int32(0) <= v2399 {
		goto L61
	} else {
		goto L591
	}
L591:
	;
	v3039 = v2399
	goto L13
L592:
	;
	v2406 = F_slice_del(m, l0)
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L8
	} else {
		goto L593
	}
L593:
	;
	if v2406 < int32(0) {
		v3039 = v2406
		goto L13
	} else {
		goto L594
	}
L594:
	;
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2410
	v2413 = v2410 - int32(1)
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2413 <= v2414 {
		goto L61
	} else {
		goto L595
	}
L595:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2416+v2413))))
	if v2418&int32(224) != int32(96) {
		goto L61
	} else {
		goto L596
	}
L596:
	;
	if int32(1)<<(uint(v2418)%32)&int32(4722696) == int32(0) {
		goto L61
	} else {
		goto L597
	}
L597:
	;
	v2431 = F_find_among_b(m, l0, int32(4342288), int32(4))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L8
	} else {
		goto L598
	}
L598:
	;
	if v2431 == int32(0) {
		goto L61
	} else {
		goto L599
	}
L599:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2435
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2437)))
	if v2435 < v2438 {
		goto L61
	} else {
		goto L600
	}
L600:
	;
	v2440 = F_slice_del(m, l0)
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L8
	} else {
		goto L601
	}
L601:
	;
	if v2440 < int32(0) {
		v3039 = v2440
		goto L13
	} else {
		goto L602
	}
L602:
	;
	if v2431 != int32(1) {
		goto L61
	} else {
		goto L603
	}
L603:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2446
	v2448 = int32(2)
	v2450 = int32(0)
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2446-v2453 < v2448 {
		v2463 = v2450
		goto L605
	} else {
		goto L606
	}
L604:
	;
	if v2463 == int32(0) {
		goto L61
	} else {
		goto L608
	}
L605:
	;
	goto L604
L606:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2459 = F_memcmp(m, v2456+v2446-v2448, int32(2242308), v2448)
	mBase = m.M
	if v2459 != 0 {
		v2463 = v2450
		goto L605
	} else {
		goto L607
	}
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2446 - v2448
	v2463 = int32(1)
	goto L605
L608:
	;
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2466
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2468)))
	if v2466 < v2469 {
		goto L61
	} else {
		goto L609
	}
L609:
	;
	v2471 = F_slice_del(m, l0)
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L8
	} else {
		goto L610
	}
L610:
	;
	if int32(0) <= v2471 {
		goto L61
	} else {
		goto L611
	}
L611:
	;
	v3039 = v2471
	goto L13
L612:
	;
	v2478 = F_slice_del(m, l0)
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L8
	} else {
		goto L613
	}
L613:
	;
	if v2478 < int32(0) {
		v3039 = v2478
		goto L13
	} else {
		goto L614
	}
L614:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2482
	v2485 = v2482 - int32(1)
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2485 <= v2486 {
		goto L61
	} else {
		goto L615
	}
L615:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2488+v2485))))
	if v2490&int32(224) != int32(96) {
		goto L61
	} else {
		goto L616
	}
L616:
	;
	if int32(1)<<(uint(v2490)%32)&int32(4198408) == int32(0) {
		goto L61
	} else {
		goto L617
	}
L617:
	;
	v2503 = F_find_among_b(m, l0, int32(4342368), int32(3))
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L8
	} else {
		goto L618
	}
L618:
	;
	if v2503 == int32(0) {
		goto L61
	} else {
		goto L619
	}
L619:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2507
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v2509)))
	if v2507 < v2510 {
		goto L61
	} else {
		goto L620
	}
L620:
	;
	v2512 = F_slice_del(m, l0)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L8
	} else {
		goto L621
	}
L621:
	;
	if int32(0) <= v2512 {
		goto L61
	} else {
		goto L622
	}
L622:
	;
	v3039 = v2512
	goto L13
L623:
	;
	v2519 = F_slice_del(m, l0)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L8
	} else {
		goto L624
	}
L624:
	;
	if v2519 < int32(0) {
		v3039 = v2519
		goto L13
	} else {
		goto L625
	}
L625:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2523
	v2525 = int32(2)
	v2527 = int32(0)
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2523-v2530 < v2525 {
		v2540 = v2527
		goto L627
	} else {
		goto L628
	}
L626:
	;
	if v2540 == int32(0) {
		goto L61
	} else {
		goto L630
	}
L627:
	;
	goto L626
L628:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2536 = F_memcmp(m, v2533+v2523-v2525, int32(2242310), v2525)
	mBase = m.M
	if v2536 != 0 {
		v2540 = v2527
		goto L627
	} else {
		goto L629
	}
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2523 - v2525
	v2540 = int32(1)
	goto L627
L630:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2543
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v2545)))
	if v2543 < v2546 {
		goto L61
	} else {
		goto L631
	}
L631:
	;
	v2548 = F_slice_del(m, l0)
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L8
	} else {
		goto L632
	}
L632:
	;
	if v2548 < int32(0) {
		v3039 = v2548
		goto L13
	} else {
		goto L633
	}
L633:
	;
	v2552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2552
	v2554 = int32(2)
	v2556 = int32(0)
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2552-v2559 < v2554 {
		v2569 = v2556
		goto L635
	} else {
		goto L636
	}
L634:
	;
	if v2569 == int32(0) {
		goto L61
	} else {
		goto L638
	}
L635:
	;
	goto L634
L636:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2565 = F_memcmp(m, v2562+v2552-v2554, int32(2242312), v2554)
	mBase = m.M
	if v2565 != 0 {
		v2569 = v2556
		goto L635
	} else {
		goto L637
	}
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2552 - v2554
	v2569 = int32(1)
	goto L635
L638:
	;
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2572
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v2574)))
	if v2572 < v2575 {
		goto L61
	} else {
		goto L639
	}
L639:
	;
	v2577 = F_slice_del(m, l0)
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L8
	} else {
		goto L640
	}
L640:
	;
	if v2577 < int32(0) {
		v3039 = v2577
		goto L13
	} else {
		goto L641
	}
L641:
	;
	goto L61
L642:
	;
	if int32(0) <= v2583 {
		goto L62
	} else {
		goto L643
	}
L643:
	;
	v3039 = v2583
	goto L13
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2589
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2591
	v2598 = F_find_among_b(m, l0, int32(4342432), int32(87))
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L8
	} else {
		goto L645
	}
L645:
	;
	if v2598 != 0 {
		goto L646
	} else {
		goto L647
	}
L646:
	;
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2600
	v2602 = F_slice_del(m, l0)
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L8
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2594
	goto L61
L649:
	;
	if v2602 < int32(0) {
		v3039 = v2602
		goto L13
	} else {
		goto L650
	}
L650:
	;
	goto L648
L651:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2780
	v2784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2780 <= v2784 {
		v2937 = v2782
		goto L686
	} else {
		goto L687
	}
L652:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2778
	v2780 = v2778
	v2782 = v2778
	goto L651
L653:
	;
	if v2742 != 0 {
		goto L652
	} else {
		goto L677
	}
L654:
	;
	v2742 = v2735
	goto L653
L655:
	;
	if v2611 <= v2631 {
		v2735 = int32(-1)
		goto L654
	} else {
		goto L657
	}
L656:
	;
	v2735 = int32(0)
	goto L654
L657:
	;
	v2648 = int32(1)
	v2649 = v2611 - v2648
	v2651 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2627+v2649))))
	v2653 = v2651 & int32(255)
	if v2649 == v2631 {
		v2708 = v2653
		v2709 = v2648
		goto L658
	} else {
		goto L659
	}
L658:
	;
	if int32(242) < v2708 {
		goto L667
	} else {
		goto L668
	}
L659:
	;
	if int32(0) <= v2651 {
		v2708 = v2653
		v2709 = v2648
		goto L658
	} else {
		goto L660
	}
L660:
	;
	v2659 = v2653 & int32(63)
	v2661 = v2611 - int32(2)
	v2663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2627+v2661))))
	v2665 = v2663 << (uint(int32(6)) % 32)
	if base.B2i32(v2661 != v2631)&base.B2i32(base.Ui32(v2663) < base.Ui32(int32(192))) == int32(0) {
		goto L661
	} else {
		goto L662
	}
L661:
	;
	v2708 = v2665&int32(1984) | v2659
	v2709 = int32(2)
	goto L658
L662:
	;
	goto L663
L663:
	;
	v2678 = v2665&int32(4032) | v2659
	v2680 = v2611 - int32(3)
	v2682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2627+v2680))))
	if base.B2i32(v2680 != v2631)&base.B2i32(base.Ui32(v2682) < base.Ui32(int32(224))) == int32(0) {
		goto L664
	} else {
		goto L665
	}
L664:
	;
	v2708 = v2682<<(uint(int32(12))%32)&int32(61440) | v2678
	v2709 = int32(3)
	goto L658
L665:
	;
	goto L666
L666:
	;
	v2700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2611+(v2627-int32(4))))))
	v2708 = v2682<<(uint(int32(12))%32)&int32(258048) | v2700&int32(7)<<(uint(int32(18))%32) | v2678
	v2709 = int32(4)
	goto L658
L667:
	;
	v2742 = v2709
	goto L653
L668:
	;
	goto L669
L669:
	;
	v2713 = v2708 - int32(97)
	if v2713 < int32(0) {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v2742 = v2709
	goto L653
L671:
	;
	goto L672
L672:
	;
	v2719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2713)>>(uint(int32(3))%32)))+uint32(_consts[1309]))))
	if int32(base.Ui32(v2719)>>(uint(v2713&int32(7))%32))&int32(1) == int32(0) {
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v2742 = v2709
	goto L653
L674:
	;
	goto L675
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2611 - v2709
	goto L676
L676:
	;
	goto L656
L677:
	;
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2743
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2745)+8))
	if v2743 < v2746 {
		goto L652
	} else {
		goto L678
	}
L678:
	;
	v2748 = F_slice_del(m, l0)
	mBase = m.M
	v2749 = m.ExcPending
	if v2749 != 0 {
		goto L8
	} else {
		goto L679
	}
L679:
	;
	if v2748 < int32(0) {
		v3039 = v2748
		goto L13
	} else {
		goto L680
	}
L680:
	;
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2752
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2752 <= v2754 {
		goto L652
	} else {
		goto L681
	}
L681:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2756+v2752-int32(1)))))
	if v2760 != int32(105) {
		goto L652
	} else {
		goto L682
	}
L682:
	;
	v2764 = v2752 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2764
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2764
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v2767)+8))
	if v2752 <= v2768 {
		goto L652
	} else {
		goto L683
	}
L683:
	;
	v2770 = F_slice_del(m, l0)
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L8
	} else {
		goto L684
	}
L684:
	;
	if v2770 < int32(0) {
		v3039 = v2770
		goto L13
	} else {
		goto L685
	}
L685:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2780 = v2774
	v2782 = v2775
	goto L651
L686:
	;
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2938
	v2941 = v2938
	v2942 = v2937
	goto L718
L687:
	;
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2786+v2780-int32(1)))))
	if v2790 != int32(104) {
		v2937 = v2782
		goto L686
	} else {
		goto L688
	}
L688:
	;
	v2794 = v2780 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2794
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L692
L689:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2937 = v2935
	goto L686
L690:
	;
	if v2925 != 0 {
		goto L689
	} else {
		goto L714
	}
L691:
	;
	v2925 = v2918
	goto L690
L692:
	;
	if v2794 <= v2814 {
		v2918 = int32(-1)
		goto L691
	} else {
		goto L694
	}
L693:
	;
	v2918 = int32(0)
	goto L691
L694:
	;
	v2831 = int32(1)
	v2832 = v2794 - v2831
	v2834 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2810+v2832))))
	v2836 = v2834 & int32(255)
	if v2832 == v2814 {
		v2891 = v2836
		v2892 = v2831
		goto L695
	} else {
		goto L696
	}
L695:
	;
	if int32(103) < v2891 {
		goto L704
	} else {
		goto L705
	}
L696:
	;
	if int32(0) <= v2834 {
		v2891 = v2836
		v2892 = v2831
		goto L695
	} else {
		goto L697
	}
L697:
	;
	v2842 = v2836 & int32(63)
	v2844 = v2794 - int32(2)
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2810+v2844))))
	v2848 = v2846 << (uint(int32(6)) % 32)
	if base.B2i32(v2844 != v2814)&base.B2i32(base.Ui32(v2846) < base.Ui32(int32(192))) == int32(0) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v2891 = v2848&int32(1984) | v2842
	v2892 = int32(2)
	goto L695
L699:
	;
	goto L700
L700:
	;
	v2861 = v2848&int32(4032) | v2842
	v2863 = v2794 - int32(3)
	v2865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2810+v2863))))
	if base.B2i32(v2863 != v2814)&base.B2i32(base.Ui32(v2865) < base.Ui32(int32(224))) == int32(0) {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	v2891 = v2865<<(uint(int32(12))%32)&int32(61440) | v2861
	v2892 = int32(3)
	goto L695
L702:
	;
	goto L703
L703:
	;
	v2883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794+(v2810-int32(4))))))
	v2891 = v2865<<(uint(int32(12))%32)&int32(258048) | v2883&int32(7)<<(uint(int32(18))%32) | v2861
	v2892 = int32(4)
	goto L695
L704:
	;
	v2925 = v2892
	goto L690
L705:
	;
	goto L706
L706:
	;
	v2896 = v2891 - int32(99)
	if v2896 < int32(0) {
		goto L707
	} else {
		goto L708
	}
L707:
	;
	v2925 = v2892
	goto L690
L708:
	;
	goto L709
L709:
	;
	v2902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2896)>>(uint(int32(3))%32)))+uint32(_consts[1310]))))
	if int32(base.Ui32(v2902)>>(uint(v2896&int32(7))%32))&int32(1) == int32(0) {
		goto L710
	} else {
		goto L711
	}
L710:
	;
	v2925 = v2892
	goto L690
L711:
	;
	goto L712
L712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2794 - v2892
	goto L713
L713:
	;
	goto L693
L714:
	;
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v2926)+8))
	v2928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2928 < v2927 {
		goto L689
	} else {
		goto L715
	}
L715:
	;
	v2930 = F_slice_del(m, l0)
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L8
	} else {
		goto L716
	}
L716:
	;
	if v2930 < int32(0) {
		v3039 = v2930
		goto L13
	} else {
		goto L717
	}
L717:
	;
	goto L689
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2941
	if v2942 <= v2941 {
		goto L723
	} else {
		goto L724
	}
L719:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2938
	v3039 = int32(1)
	goto L13
L720:
	;
	goto L719
L721:
	;
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2941 = v3035
	v2942 = v3034
	goto L718
L722:
	;
	v2976 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L736
L723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2941
	v2974 = v2941
	v2975 = v2942
	goto L722
L724:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2947+v2941))))
	switch v2949 - int32(73) {
	case 0, 12:
		goto L725
	default:
		goto L723
	}
L725:
	;
	v2954 = F_find_among(m, l0, int32(4344176), int32(3))
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L8
	} else {
		goto L726
	}
L726:
	;
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2956
	switch v2954 - int32(1) {
	case 0:
		goto L728
	case 1:
		goto L727
	case 2:
		goto L729
	default:
		goto L721
	}
L727:
	;
	v2969 = F_slice_from_s(m, l0, int32(1), int32(2242965))
	mBase = m.M
	v2970 = m.ExcPending
	if v2970 != 0 {
		goto L8
	} else {
		goto L732
	}
L728:
	;
	v2963 = F_slice_from_s(m, l0, int32(1), int32(2242964))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L8
	} else {
		goto L730
	}
L729:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2974 = v2956
	v2975 = v2960
	goto L722
L730:
	;
	if int32(0) <= v2963 {
		goto L721
	} else {
		goto L731
	}
L731:
	;
	v3039 = v2963
	goto L13
L732:
	;
	if int32(0) <= v2969 {
		goto L721
	} else {
		goto L733
	}
L733:
	;
	v3039 = v2969
	goto L13
L734:
	;
	if v3028 < int32(0) {
		goto L720
	} else {
		goto L754
	}
L736:
	;
	goto L737
L737:
	;
	goto L738
L738:
	;
	v2983 = v2974
	v2985 = int32(1)
	goto L741
L740:
	;
	v3028 = v3013
	goto L734
L741:
	;
	if v2975 <= v2983 {
		goto L743
	} else {
		goto L744
	}
L742:
	;
	goto L740
L743:
	;
	v3028 = int32(-1)
	goto L734
L744:
	;
	goto L745
L745:
	;
	v2990 = v2983 + int32(1)
	v2992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2976+v2983))))
	if base.Ui32(v2992) < base.Ui32(int32(192)) {
		v3013 = v2990
		goto L746
	} else {
		goto L747
	}
L746:
	;
	v3014 = int32(1)
	if v3014 < v2985 {
		v2983 = v3013
		v2985 = v2985 - v3014
		goto L741
	} else {
		goto L753
	}
L747:
	;
	if v2975 <= v2990 {
		v3013 = v2990
		goto L746
	} else {
		goto L748
	}
L748:
	;
	v2999 = v2990
	goto L749
L749:
	;
	v3002 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2976+v2999))))
	if int32(-65) < v3002 {
		v3013 = v2999
		goto L746
	} else {
		goto L751
	}
L750:
	;
	v3013 = v2975
	goto L746
L751:
	;
	v3006 = v2999 + int32(1)
	if v3006 != v2975 {
		v2999 = v3006
		goto L749
	} else {
		goto L752
	}
L752:
	;
	goto L750
L753:
	;
	goto L742
L754:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3028
	goto L721
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 float64
	_ = v97
	var v108 float64
	_ = v108
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	v20 = m.G0
	v22 = v20 - int32(80)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v24 != 0 {
		v25 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v25
		v28 = v22 - int32(-64)
		*(*int64)(unsafe.Add(mBase, uint32(v28))) = v25
		*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v25
		v34 = v22 + int32(48)
		*(*int64)(unsafe.Add(mBase, uint32(v34))) = v25
		*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v25
		*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v25
		v42 = v22 + int32(24)
		*(*int64)(unsafe.Add(mBase, uint32(v42))) = v25
		*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v25
		F_genericcostestimate(m, l0, l1, l2, v22+int32(16))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
			v54 = F_index_open(m, v52, int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				F_IvfflatGetMetaPageInfo(m, v54, v22+int32(12), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_relation_close(m, v54, int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, _consts[1463]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
						F_get_tablespace_page_costs(m, v68, int32(0), v22)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							v72 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v22)))
							v77 = base.F64_sub(v75, v76)
							v79 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
							v80 = base.F64_add(base.F64_mul(base.F64_mul(v72, float64(-0.5)), v77), v79)
							v81 = float64(1)
							v84 = base.F64_div(base.F64_convert_i32_s(v65), base.F64_convert_i32_s(v66))
							if base.F64_gt(v84, v81) != 0 {
								v87 = v81
							} else {
								v87 = v84
							}
							v88 = base.F64_mul(v80, v87)
							if base.F64_lt(v87, float64(0.5)) == int32(0) {
								v108 = v88
							} else {
								v93 = base.F64_mul(v72, v87)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+116))
								v97 = base.F64_convert_i32_u(v96)
								if base.F64_gt(v93, v97) == int32(0) {
									v108 = v88
								} else {
									v108 = base.F64_add(base.F64_mul(base.F64_sub(v97, v93), v76), base.F64_add(base.F64_mul(base.F64_mul(v93, float64(-0.5)), v77), v88))
								}
							}
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = v108
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v80
							v113 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
							*(*float64)(unsafe.Add(mBase, uint32(l5))) = v113
							v115 = *(*float64)(unsafe.Add(mBase, uint32(v22)+40))
							*(*float64)(unsafe.Add(mBase, uint32(l6))) = v115
							*(*float64)(unsafe.Add(mBase, uint32(l7))) = v72
							m.G0 = v22 + int32(80)
							return
						}
					}
				}
			}
		}
	} else {
		v118 = int64(9218868437227405312)
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v118
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v118
		v122 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l5))) = v122
		*(*int64)(unsafe.Add(mBase, uint32(l6))) = v122
		*(*int64)(unsafe.Add(mBase, uint32(l7))) = v122
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(2)
		m.G0 = v22 + int32(80)
		return
	}
}
