package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecEvalSysVar(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v5 - int32(1) {
	case 0:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v13&int32(8) != 0 {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v55 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			switch v17 + int32(6) {
			case 0:
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v20)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
				v31 = v22
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
				if v35 != int32(1) {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			default:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
				v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l2, v17, v16)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
					if v35 != int32(1) {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
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
			case 5:
				v23 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v23)
				v31 = l2 + int32(28)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
				if v35 != int32(1) {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
		}
	case 1:
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v8&int32(16) == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			switch v17 + int32(6) {
			case 0:
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v20)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
				v31 = v22
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
				if v35 != int32(1) {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			default:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
				v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l2, v17, v16)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
					if v35 != int32(1) {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
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
			case 5:
				v23 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v23)
				v31 = l2 + int32(28)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
				if v35 != int32(1) {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v55 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
			return
		}
	default:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		switch v17 + int32(6) {
		case 0:
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v20)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
			v31 = v22
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
			if v35 != int32(1) {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		default:
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l2, v17, v16)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = v29
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
				if v35 != int32(1) {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
		case 5:
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v23)
			v31 = l2 + int32(28)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
			if v35 != int32(1) {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ExecEvalSysVar_0), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecEvalSysVar_1), int32(_a_F_ExecEvalSysVar_2), int32(_a_F_ExecEvalSysVar_3))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
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
	}
}
func F_GetSysCacheOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l0 << (uint(int32(2)) % 32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_GetSysCacheOid[0])))
	v16 = F_SearchCatCache(m, v15, l1, l2, l3, l4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v32 = int32(0)
			m.G0 = v9 + int32(16)
			return v32
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_GetSysCacheOid[0])))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			v28 = F_heap_getattr_1(m, v16, int32(1), v25, v9+int32(15))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v16)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = v28
					m.G0 = v9 + int32(16)
					return v32
				}
			}
		}
	}
}
func F_SearchSysCacheCopyAttNum(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13859(m, l0, l1, int32(_a_F_SearchSysCacheCopyAttNum_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_SearchSysCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(l0) <= base.Ui32(int32(84)) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCacheList[0])))
		if v16 != 0 {
			v33 = F_SearchCatCacheList(m, v16, l1, l2, l3, l4)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v33
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_SearchSysCacheList_0), v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_SearchSysCacheList_1), int32(683), int32(_a_F_SearchSysCacheList_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			F_errmsg_internal(m, int32(_a_F_SearchSysCacheList_0), v10)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_SearchSysCacheList_1), int32(683), int32(_a_F_SearchSysCacheList_2))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
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
func F_StartSysLogger(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	v8 = F_AssignPostmasterChildSlot(m, int32(17))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[0])) = v8
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[1]))
	if v17 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L100
	}
L6:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v243
	if v243 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L92
	}
L8:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L88
	}
L9:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L84
	}
L10:
	;
	v21 = F_pipe(m, int32(_a_F_StartSysLogger_0))
	mBase = m.M
	if v21 < int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[2]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[3]))
	v28 = F_mkdir(m, v25, v27)
	mBase = m.M
	goto L14
L13:
	;
	goto L12
L14:
	;
	v30 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_StartSysLogger[4])) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v30
	v34 = F_palloc(m, int32(1024))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v37
	v43 = F_pg_snprintf(m, v34, int32(1024), int32(_a_F_StartSysLogger_4), v14+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v45 = F_strlen(m, v34)
	mBase = m.M
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[5]))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[6]))
	v55 = F_pg_localtime(m, v14+int32(32), v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v57 = F_pg_strftime(m, v34+v45, int32(1024)-v45, v50, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v62 = F_logfile_open(m, v34, int32(_a_F_StartSysLogger_5), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[7])) = v62
	F_pfree(m, v34)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[8]))
	if v68&int32(8) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v73 = *(*int64)(unsafe.Add(mBase, _c_F_StartSysLogger[4]))
	v75 = F_logfile_getname(m, v73, int32(_a_F_StartSysLogger_6))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v87 = v68
	goto L23
L23:
	;
	if v87&int32(16) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v79 = F_logfile_open(m, v75, int32(_a_F_StartSysLogger_5), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[9])) = v79
	F_pfree(m, v75)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[8]))
	v87 = v85
	goto L23
L27:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_StartSysLogger[4]))
	v94 = F_logfile_getname(m, v92, int32(_a_F_StartSysLogger_7))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[7]))
	if v105 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v98 = F_logfile_open(m, v94, int32(_a_F_StartSysLogger_5), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[10])) = v98
	F_pfree(m, v94)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+60))
	if v106 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v115 = int32(-1)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v115
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[9]))
	if v118 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v115 = v113
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[11])) = int32(8)
	v113 = int32(-1)
	goto L39
L38:
	;
	v113 = v106
	goto L39
L39:
	;
	goto L36
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+60))
	if v119 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v128 = int32(-1)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v128
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[10]))
	if v131 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v128 = v126
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[11])) = int32(8)
	v126 = int32(-1)
	goto L46
L45:
	;
	v126 = v119
	goto L46
L46:
	;
	goto L43
L47:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+60))
	if v132 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v141 = int32(-1)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v141
	v148 = F_postmaster_child_launch(m, int32(17), v11, v14+int32(32), int32(12), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L55
	}
L50:
	;
	v141 = v139
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[11])) = int32(8)
	v139 = int32(-1)
	goto L53
L52:
	;
	v139 = v132
	goto L53
L53:
	;
	goto L50
L54:
	;
	m.G0 = v14 + int32(48)
	goto L6
L55:
	;
	if v148 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = int32(0)
	v155 = F_errstart(m, int32(15), v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartSysLogger[12])))
	if v169 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	if v155 == int32(0) {
		v243 = v152
		goto L54
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_StartSysLogger_8), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_StartSysLogger_2), int32(711), int32(_a_F_StartSysLogger_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v243 = v152
	goto L54
L63:
	;
	v174 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[7]))
	v222 = F_fclose(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L77
	}
L66:
	;
	if v174 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_errmsg(m, int32(_a_F_StartSysLogger_9), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[13]))
	v193 = F_fflush(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v181
	F_errhint(m, int32(_a_F_StartSysLogger_10), v14)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_StartSysLogger_2), int32(732), int32(_a_F_StartSysLogger_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[14]))
	v198 = F_dup2(m, v196, int32(1))
	mBase = m.M
	if v198 < int32(0) {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[15]))
	v203 = F_fflush(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[14]))
	v208 = F_dup2(m, v206, int32(2))
	mBase = m.M
	if v208 < int32(0) {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v211 = int32(_a_F_StartSysLogger_13)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[14]))
	v213 = F_close(m, v212)
	mBase = m.M
	v215 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartSysLogger[12])) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[14])) = int32(-1)
	goto L65
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[7])) = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[9]))
	if v228 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v229 = F_fclose(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[10]))
	if v235 == int32(0) {
		v243 = v148
		goto L54
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[9])) = int32(0)
	goto L80
L82:
	;
	v238 = F_fclose(m, v235)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[10])) = int32(0)
	v243 = v148
	goto L54
L84:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_StartSysLogger_1), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_StartSysLogger_2), int32(626), int32(_a_F_StartSysLogger_3))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_StartSysLogger_11), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_StartSysLogger_2), int32(739), int32(_a_F_StartSysLogger_3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_StartSysLogger_12), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_StartSysLogger_2), int32(744), int32(_a_F_StartSysLogger_3))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v298 = F_ReleasePostmasterChildSlot(m, v294)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	return
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSysLogger[0])) = int32(0)
	goto L98
L100:
	;
	F_errmsg_internal(m, int32(_a_F_StartSysLogger_14), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_StartSysLogger_15), int32(4012), int32(_a_F_StartSysLogger_16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
