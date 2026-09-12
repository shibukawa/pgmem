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
						F_errmsg_internal(m, int32(85130), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
							F_errmsg_internal(m, int32(85130), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
						F_errmsg_internal(m, int32(85130), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
						F_errmsg_internal(m, int32(85130), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
							F_errmsg_internal(m, int32(85130), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
						F_errmsg_internal(m, int32(85130), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
					F_errmsg_internal(m, int32(85130), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
						F_errmsg_internal(m, int32(85130), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
					F_errmsg_internal(m, int32(85130), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(490751), int32(5608), int32(228511))
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
				F_errmsg_internal(m, int32(482987), v10)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494170), int32(683), int32(75902))
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
			F_errmsg_internal(m, int32(482987), v10)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494170), int32(683), int32(75902))
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
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
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
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L81
	}
L6:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[455]))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v201
	if v201 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L73
	}
L8:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L69
	}
L9:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L65
	}
L10:
	;
	v21 = F_pipe(m, int32(4095420))
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
	v43 = F_pg_snprintf(m, v34, int32(1024), int32(550183), v14+int32(16))
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
	v50 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	v54 = *(*int32)(unsafe.Add(mBase, _consts[144]))
	v55 = F_pg_localtime(m, v14+int32(24), v54)
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
	v62 = F_logfile_open(m, v34, int32(502219), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[460])) = v62
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
	v68 = *(*int32)(unsafe.Add(mBase, _consts[461]))
	if v68&int32(8) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v73 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	v75 = F_logfile_getname(m, v73, int32(32404))
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
	v79 = F_logfile_open(m, v75, int32(502219), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[462])) = v79
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
	v85 = *(*int32)(unsafe.Add(mBase, _consts[461]))
	v87 = v85
	goto L23
L27:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	v94 = F_logfile_getname(m, v92, int32(243125))
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
	v104 = int32(0)
	v109 = F_postmaster_child_launch(m, int32(17), v11, v104, v104, v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v98 = F_logfile_open(m, v94, int32(502219), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[463])) = v98
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
	m.G0 = v14 + int32(32)
	goto L6
L34:
	;
	if v109 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v115 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[464])))
	if v129 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	if v115 == int32(0) {
		v201 = v104
		goto L33
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(290920), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(490504), int32(711), int32(82452))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v201 = v104
	goto L33
L42:
	;
	v134 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v182 = F_fclose(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L56
	}
L45:
	;
	if v134 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_errmsg(m, int32(127561), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	v153 = F_fflush(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v141
	F_errhint(m, int32(643892), v14)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(490504), int32(732), int32(82452))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v158 = F_dup2(m, v156, int32(1))
	mBase = m.M
	if v158 < int32(0) {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	v163 = F_fflush(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v168 = F_dup2(m, v166, int32(2))
	mBase = m.M
	if v168 < int32(0) {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v171 = int32(4095424)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v173 = F_close(m, v172)
	mBase = m.M
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[464])) = uint8(v175)
	*(*int32)(unsafe.Add(mBase, _consts[465])) = int32(-1)
	goto L44
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[460])) = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	if v188 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v189 = F_fclose(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	if v195 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[462])) = int32(0)
	goto L59
L61:
	;
	v196 = F_fclose(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v201 = v109
	goto L33
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[463])) = int32(0)
	goto L63
L65:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(291264), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(490504), int32(626), int32(82452))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(289622), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(490504), int32(739), int32(82452))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(290774), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(490504), int32(744), int32(82452))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v255 = F_ReleasePostmasterChildSlot(m, v251)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	return
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[455])) = int32(0)
	goto L79
L81:
	;
	F_errmsg_internal(m, int32(221898), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(490240), int32(4012), int32(221947))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
