package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_avg_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 == int32(0) {
		v35 = int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		switch v10 - int32(429) {
		case 0:
			v35 = int32(1)
		case 1:
			v35 = int32(2)
		default:
			v35 = int32(0)
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v35 != 0 {
		v37 = F_pg_detoast_datum(m, v36)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v43 = v37
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 + int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 + v4
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(26664), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525301), int32(6822), int32(301335))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(26664), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525301), int32(6822), int32(301335))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
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
		v41 = F_pg_detoast_datum_copy(m, v36)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 + int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 + v4
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(26664), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525301), int32(6822), int32(301335))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(26664), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525301), int32(6822), int32(301335))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
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
	}
}
func F_int4_avg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(0) {
		v34 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		switch v9 - int32(429) {
		case 0:
			v34 = int32(1)
		case 1:
			v34 = int32(2)
		default:
			v34 = int32(0)
		}
	}
	if v34 != 0 {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v36 = F_pg_detoast_datum(m, v35)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v41 = F_pg_detoast_datum(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
				if v43 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(26664), int32(0))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525301), int32(6847), int32(392310))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					if v44&int32(-4) != int32(160) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(26664), int32(0))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(525301), int32(6847), int32(392310))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
						if v49 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(26664), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(525301), int32(6851), int32(392310))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
							if v50&int32(-4) != int32(160) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(26664), int32(0))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(525301), int32(6851), int32(392310))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
								v56 = int32(3)
								v58 = int32(23)
								v60 = int32(-8)
								v62 = v36 + (v55<<(uint(v56)%32)+v58)&v60
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
								v71 = v41 + (v64<<(uint(v56)%32)+v58)&v60
								v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
								*(*int64)(unsafe.Add(mBase, uint32(v62))) = v63 + v72
								v75 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = v75 + v76
								return v36
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66939), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(525301), int32(6840), int32(392310))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
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
func F_int4_dist(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = v7 - v4
	if base.B2i32(base.B2i32(v2 < v4)^base.B2i32(v8 < v7) == v2)&base.B2i32(v8 != int32(-2147483648)) == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(421613), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(525896), int32(105), int32(81947))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
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
		v41 = v8 >> (uint(int32(31)) % 32)
		return v8 ^ v41 - v41
	}
}
