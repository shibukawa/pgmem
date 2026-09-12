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
						F_errmsg_internal(m, int32(85079), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
							F_errmsg_internal(m, int32(85079), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
						F_errmsg_internal(m, int32(85079), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
						F_errmsg_internal(m, int32(85079), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
							F_errmsg_internal(m, int32(85079), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
						F_errmsg_internal(m, int32(85079), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
					F_errmsg_internal(m, int32(85079), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
						F_errmsg_internal(m, int32(85079), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
					F_errmsg_internal(m, int32(85079), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(490299), int32(5608), int32(228273))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = l0 << (uint(int32(2)) % 32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1140])))
	v17 = F_SearchCatCache(m, v16, l1, l2, l3, l4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			v33 = int32(0)
			m.G0 = v10 + int32(16)
			return v33
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1140])))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			v29 = F_heap_getattr_1(m, v17, int32(1), v26, v10+int32(15))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v17)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = v29
					m.G0 = v10 + int32(16)
					return v33
				}
			}
		}
	}
}
func F_SearchSysCacheCopyAttNum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1141]))
	v6 = F_SearchCatCache2(m, v5, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11)+91)))
			if v13 == int32(0) {
				v16 = F_heap_copytuple(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = v16
					F_ReleaseCatCache(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = v18
						return v22
					}
				}
			} else {
				v18 = v3
				F_ReleaseCatCache(m, v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v22 = v18
					return v22
				}
			}
		} else {
			v22 = v3
			return v22
		}
	}
}
func F_SearchSysCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(l0) <= base.Ui32(int32(84)) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1140])))
		if v18 != 0 {
			v35 = F_SearchCatCacheList(m, v18, l1, l2, l3, l4)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v35
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(482535), v10)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493702), int32(683), int32(75851))
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			F_errmsg_internal(m, int32(482535), v10)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493702), int32(683), int32(75851))
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
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
	*(*int32)(unsafe.Add(mBase, _consts[455])) = v8
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[456]))
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
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L98
	}
L6:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[455]))
	*(*int32)(unsafe.Add(mBase, uint32(v307))) = v257
	if v257 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L90
	}
L8:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L86
	}
L9:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L82
	}
L10:
	;
	v21 = F_pipe(m, int32(4094140))
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	v27 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	v28 = F_mkdir(m, v25, v27)
	mBase = m.M
	goto L14
L13:
	;
	goto L12
L14:
	;
	v30 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[458])) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v30
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
	v37 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v37
	v43 = F_pg_snprintf(m, v34, int32(1024), int32(549285), v14+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v34&int32(3) == int32(0) {
		v68 = v34
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[144]))
	v111 = F_pg_localtime(m, v14+int32(24), v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L18:
	;
	v101 = v93 - v34
	goto L17
L19:
	;
	v72 = v68
	goto L28
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v52 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v101 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v57 = v34
	goto L24
L24:
	;
	v61 = v57 + int32(1)
	if v61&int32(3) == int32(0) {
		v68 = v61
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v93 = v61
	goto L18
L26:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v66 != 0 {
		v57 = v61
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = int32(-2139062144)
	if (int32(16843008)-v78|v78)&v81 == v81 {
		v72 = v72 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v87 = v72
	goto L31
L30:
	;
	goto L29
L31:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 != 0 {
		v87 = v87 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v93 = v87
	goto L18
L33:
	;
	goto L32
L34:
	;
	v113 = F_pg_strftime(m, v34+v101, int32(1024)-v101, v106, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v118 = F_logfile_open(m, v34, int32(501751), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[460])) = v118
	F_pfree(m, v34)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[461]))
	if v124&int32(8) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v129 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	v131 = F_logfile_getname(m, v129, int32(32387))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v143 = v124
	goto L40
L40:
	;
	if v143&int32(16) != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v135 = F_logfile_open(m, v131, int32(501751), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[462])) = v135
	F_pfree(m, v131)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[461]))
	v143 = v141
	goto L40
L44:
	;
	v148 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	v150 = F_logfile_getname(m, v148, int32(242887))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v160 = int32(0)
	v165 = F_postmaster_child_launch(m, int32(17), v11, v160, v160, v160)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L51
	}
L47:
	;
	v154 = F_logfile_open(m, v150, int32(501751), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[463])) = v154
	F_pfree(m, v150)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	m.G0 = v14 + int32(32)
	goto L6
L51:
	;
	if v165 == int32(-1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v171 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, _consts[464])))
	if v185 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	if v171 == int32(0) {
		v257 = v160
		goto L50
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(290642), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(490052), int32(711), int32(82401))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v257 = v160
	goto L50
L59:
	;
	v190 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v238 = F_fclose(m, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L73
	}
L62:
	;
	if v190 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_errmsg(m, int32(127417), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	v209 = F_fflush(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v197
	F_errhint(m, int32(642994), v14)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(490052), int32(732), int32(82401))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v214 = F_dup2(m, v212, int32(1))
	mBase = m.M
	if v214 < int32(0) {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	v219 = F_fflush(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v224 = F_dup2(m, v222, int32(2))
	mBase = m.M
	if v224 < int32(0) {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v227 = int32(4094144)
	v228 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v229 = F_close(m, v228)
	mBase = m.M
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[464])) = uint8(v231)
	*(*int32)(unsafe.Add(mBase, _consts[465])) = int32(-1)
	goto L61
L73:
	;
	*(*int32)(unsafe.Add(mBase, _consts[460])) = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	if v244 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v245 = F_fclose(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	if v251 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[462])) = int32(0)
	goto L76
L78:
	;
	v252 = F_fclose(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v257 = v165
	goto L50
L81:
	;
	*(*int32)(unsafe.Add(mBase, _consts[463])) = int32(0)
	goto L80
L82:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(290986), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(490052), int32(626), int32(82401))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(289344), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(490052), int32(739), int32(82401))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(290496), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(490052), int32(744), int32(82401))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v311 = F_ReleasePostmasterChildSlot(m, v307)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	return
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[455])) = int32(0)
	goto L96
L98:
	;
	F_errmsg_internal(m, int32(221660), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(489788), int32(4012), int32(221709))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
