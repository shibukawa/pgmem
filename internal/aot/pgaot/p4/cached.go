package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SaveCachedPlan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v4 != int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		if v7 == int32(0) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v25 = *(*int32)(unsafe.Add(mBase, _consts[373]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			if v29 != v25 {
				if v29 == int32(0) {
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
					if v34 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v33
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v33
					}
					if v33 == int32(0) {
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v39
					}
				}
				if v25 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v46
					if v46 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v23
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v23
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
				}
			} else {
			}
			v60 = *(*int32)(unsafe.Add(mBase, _consts[459]))
			if v60 != 0 {
				v62 = *(*int32)(unsafe.Add(mBase, _consts[1369]))
				v67 = v62
			} else {
				v64 = int32(4062704)
				*(*int32)(unsafe.Add(mBase, _consts[459])) = v64
				v67 = v64
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v67
			v69 = int32(4062704)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v69
			v72 = l0 + int32(100)
			*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v72
			*(*int32)(unsafe.Add(mBase, _consts[1369])) = v72
			v76 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v76)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
			v14 = v12 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v14
			if v14 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v25 = *(*int32)(unsafe.Add(mBase, _consts[373]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				if v29 != v25 {
					if v29 == int32(0) {
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
						if v34 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v33
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v33
						}
						if v33 == int32(0) {
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v39
						}
					}
					if v25 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v46
						if v46 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v23
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v23
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
					}
				} else {
				}
				v60 = *(*int32)(unsafe.Add(mBase, _consts[459]))
				if v60 != 0 {
					v62 = *(*int32)(unsafe.Add(mBase, _consts[1369]))
					v67 = v62
				} else {
					v64 = int32(4062704)
					*(*int32)(unsafe.Add(mBase, _consts[459])) = v64
					v67 = v64
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v67
				v69 = int32(4062704)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v69
				v72 = l0 + int32(100)
				*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v72
				*(*int32)(unsafe.Add(mBase, _consts[1369])) = v72
				v76 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v76)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)))
				if v18 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v25 = *(*int32)(unsafe.Add(mBase, _consts[373]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					if v29 != v25 {
						if v29 == int32(0) {
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
							if v34 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v33
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v33
							}
							if v33 == int32(0) {
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v39
							}
						}
						if v25 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v46
							if v46 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v23
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v23
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
						}
					} else {
					}
					v60 = *(*int32)(unsafe.Add(mBase, _consts[459]))
					if v60 != 0 {
						v62 = *(*int32)(unsafe.Add(mBase, _consts[1369]))
						v67 = v62
					} else {
						v64 = int32(4062704)
						*(*int32)(unsafe.Add(mBase, _consts[459])) = v64
						v67 = v64
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v67
					v69 = int32(4062704)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v69
					v72 = l0 + int32(100)
					*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v72
					*(*int32)(unsafe.Add(mBase, _consts[1369])) = v72
					v76 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v76)
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					F_MemoryContextDelete(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v25 = *(*int32)(unsafe.Add(mBase, _consts[373]))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
						if v29 != v25 {
							if v29 == int32(0) {
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
								if v34 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v33
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v33
								}
								if v33 == int32(0) {
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v39
								}
							}
							if v25 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v46
								if v46 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v23
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v23
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
							}
						} else {
						}
						v60 = *(*int32)(unsafe.Add(mBase, _consts[459]))
						if v60 != 0 {
							v62 = *(*int32)(unsafe.Add(mBase, _consts[1369]))
							v67 = v62
						} else {
							v64 = int32(4062704)
							*(*int32)(unsafe.Add(mBase, _consts[459])) = v64
							v67 = v64
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v67
						v69 = int32(4062704)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v69
						v72 = l0 + int32(100)
						*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v72
						*(*int32)(unsafe.Add(mBase, _consts[1369])) = v72
						v76 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v76)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(269231), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				F_errfinish(m, int32(475827), int32(539), int32(269512))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
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
func F_get_cached_rowtype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l0 != int32(2249) {
		if v12 == int32(0) {
			v24 = F_lookup_type_cache(m, l0, int32(256))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+188))
				if v28 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							v79 = F_format_type_be(m, l0)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v79
								F_errmsg(m, int32(332183), v10)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472584), int32(2504), int32(347904))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
					*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v32
					if l3 == int32(0) {
						v38 = v24
					} else {
						v36 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v36)
						v38 = v24
					}
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+188))
					v66 = v40
					m.G0 = v10 + int32(16)
					return v66
				}
			}
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
			if v17 == int64(0) {
				v24 = F_lookup_type_cache(m, l0, int32(256))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+188))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = F_format_type_be(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v79
									F_errmsg(m, int32(332183), v10)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472584), int32(2504), int32(347904))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
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
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
						v32 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
						*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v32
						if l3 == int32(0) {
							v38 = v24
						} else {
							v36 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v36)
							v38 = v24
						}
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+188))
						v66 = v40
						m.G0 = v10 + int32(16)
						return v66
					}
				}
			} else {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v12)+192))
				if v20 == v17 {
					v38 = v12
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+188))
					v66 = v40
					m.G0 = v10 + int32(16)
					return v66
				} else {
					v24 = F_lookup_type_cache(m, l0, int32(256))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+188))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = F_format_type_be(m, l0)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v79
										F_errmsg(m, int32(332183), v10)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472584), int32(2504), int32(347904))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
							v32 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
							*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v32
							if l3 == int32(0) {
								v38 = v24
							} else {
								v36 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v36)
								v38 = v24
							}
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+188))
							v66 = v40
							m.G0 = v10 + int32(16)
							return v66
						}
					}
				}
			}
		}
	} else {
		if v12 == int32(0) {
			v52 = F_lookup_rowtype_tupdesc(m, int32(2249), l1)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
				if int32(0) <= v54 {
					F_DecrTupleDescRefCount(m, v52)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
						if l3 == int32(0) {
							v66 = v52
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
							v66 = v52
						}
						m.G0 = v10 + int32(16)
						return v66
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
					if l3 == int32(0) {
						v66 = v52
					} else {
						v64 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
						v66 = v52
					}
					m.G0 = v10 + int32(16)
					return v66
				}
			}
		} else {
			v43 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
			if v43 != int64(0) {
				v52 = F_lookup_rowtype_tupdesc(m, int32(2249), l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
					if int32(0) <= v54 {
						F_DecrTupleDescRefCount(m, v52)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
							if l3 == int32(0) {
								v66 = v52
							} else {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
								v66 = v52
							}
							m.G0 = v10 + int32(16)
							return v66
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
						if l3 == int32(0) {
							v66 = v52
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
							v66 = v52
						}
						m.G0 = v10 + int32(16)
						return v66
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				if v46 != int32(2249) {
					v52 = F_lookup_rowtype_tupdesc(m, int32(2249), l1)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
						if int32(0) <= v54 {
							F_DecrTupleDescRefCount(m, v52)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
								if l3 == int32(0) {
									v66 = v52
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
									v66 = v52
								}
								m.G0 = v10 + int32(16)
								return v66
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
							if l3 == int32(0) {
								v66 = v52
							} else {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
								v66 = v52
							}
							m.G0 = v10 + int32(16)
							return v66
						}
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if l1 == v49 {
						v66 = v12
						m.G0 = v10 + int32(16)
						return v66
					} else {
						v52 = F_lookup_rowtype_tupdesc(m, int32(2249), l1)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
							if int32(0) <= v54 {
								F_DecrTupleDescRefCount(m, v52)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
									if l3 == int32(0) {
										v66 = v52
									} else {
										v64 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
										v66 = v52
									}
									m.G0 = v10 + int32(16)
									return v66
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
								if l3 == int32(0) {
									v66 = v52
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
									v66 = v52
								}
								m.G0 = v10 + int32(16)
								return v66
							}
						}
					}
				}
			}
		}
	}
}
