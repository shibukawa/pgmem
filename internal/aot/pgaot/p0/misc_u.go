package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateChangedParamSet(m *base.Module, l0 int32, l1 int32) {
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+68))
	v5 = F_bms_intersect(m, v4, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v8 = F_bms_join(m, v7, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v8
			return
		}
	}
}
func F_uint64in_subr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int64
	_ = v118
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, _c_F_uint64in_subr[0])) = v4
	v19 = F_strtox_2(m, l0, v10+int32(44), v4, int64(-1))
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_uint64in_subr[0]))
	if v21 != 0 {
		v25 = base.B2i32(v21 != int32(68))
	} else {
		v25 = int32(0)
	}
	if v25 == int32(0) {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
		if v28 != l0 {
			if v21 == int32(68) {
				v53 = int64(0)
				v54 = F_errsave_start(m, l2)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					if v54 == int32(0) {
						v118 = v53
						m.G0 = v10 + int32(48)
						return v118
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
							F_errmsg(m, int32(_a_F_uint64in_subr_0), v10+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, l2, int32(_a_F_uint64in_subr_1), int32(1009), int32(_a_F_uint64in_subr_2))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int64(0)
								} else {
									v118 = v53
									m.G0 = v10 + int32(48)
									return v118
								}
							}
						}
					}
				}
			} else {
				v78 = v28
				for {
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
					if base.B2i32(base.Ui32(v80-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v80 == int32(32)) != 0 {
						v78 = v78 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				if v80 == int32(0) {
					v118 = v19
					m.G0 = v10 + int32(48)
					return v118
				} else {
					v92 = int64(0)
					v93 = F_errsave_start(m, l2)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						if v93 == int32(0) {
							v118 = v92
							m.G0 = v10 + int32(48)
							return v118
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l1
								F_errmsg(m, int32(_a_F_uint64in_subr_3), v10+int32(32))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, l2, int32(_a_F_uint64in_subr_1), int32(1025), int32(_a_F_uint64in_subr_2))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int64(0)
									} else {
										v118 = v92
										m.G0 = v10 + int32(48)
										return v118
									}
								}
							}
						}
					}
				}
			}
		} else {
			v31 = int64(0)
			v32 = F_errsave_start(m, l2)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				if v32 == int32(0) {
					v118 = v31
					m.G0 = v10 + int32(48)
					return v118
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg(m, int32(_a_F_uint64in_subr_3), v10)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, l2, int32(_a_F_uint64in_subr_1), int32(1003), int32(_a_F_uint64in_subr_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								v118 = v31
								m.G0 = v10 + int32(48)
								return v118
							}
						}
					}
				}
			}
		}
	} else {
		v31 = int64(0)
		v32 = F_errsave_start(m, l2)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			if v32 == int32(0) {
				v118 = v31
				m.G0 = v10 + int32(48)
				return v118
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg(m, int32(_a_F_uint64in_subr_3), v10)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, l2, int32(_a_F_uint64in_subr_1), int32(1003), int32(_a_F_uint64in_subr_2))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							v118 = v31
							m.G0 = v10 + int32(48)
							return v118
						}
					}
				}
			}
		}
	}
}
func F_unlink_span(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v46 = v10
		if v46 != 0 {
			v50 = int32(0)
			v53 = base.AtomicRmwOr32(m, v50, int32(_a_F_unlink_span_0), v50)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
			if v54 != v56 {
				v61 = F_LWLockAcquire(m, v55+int32(1476), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v65+int32(1476))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
							v76 = l0 + v73*int32(20)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
							if v77 != 0 {
								v81 = v77
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
								return
							} else {
								v78 = F_get_segment_by_index(m, l0, v73)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									v81 = v80
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
									return
								}
							}
						}
					}
				}
			} else {
				v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
				v76 = l0 + v73*int32(20)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
				if v77 != 0 {
					v81 = v77
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
					return
				} else {
					v78 = F_get_segment_by_index(m, l0, v73)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
						v81 = v80
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
						return
					}
				}
			}
		} else {
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v86 = int32(0)
			v89 = base.AtomicRmwOr32(m, v86, int32(_a_F_unlink_span_0), v86)
			v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+1468))
			if v90 != v92 {
				v97 = F_LWLockAcquire(m, v91+int32(1476), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v101+int32(1476))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
							v112 = l0 + v109*int32(20)
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
							if v113 != 0 {
								v117 = v113
								v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
								return
							} else {
								v114 = F_get_segment_by_index(m, l0, v109)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
									v117 = v116
									v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
									return
								}
							}
						}
					}
				}
			} else {
				v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
				v112 = l0 + v109*int32(20)
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
				if v113 != 0 {
					v117 = v113
					v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
					return
				} else {
					v114 = F_get_segment_by_index(m, l0, v109)
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
						v117 = v116
						v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
						return
					}
				}
			}
		}
	} else {
		v11 = int32(0)
		v14 = base.AtomicRmwOr32(m, v11, int32(_a_F_unlink_span_0), v11)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1468))
		if v15 != v17 {
			v22 = F_LWLockAcquire(m, v16+int32(1476), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v26+int32(1476))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v34 = int32(base.Ui32(v7) >> (uint(int32(27)) % 32))
						v37 = l0 + v34*int32(20)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
						if v38 != 0 {
							v42 = v38
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v42+v7&int32(134217727))+4)) = v44
							v46 = v44
							if v46 != 0 {
								v50 = int32(0)
								v53 = base.AtomicRmwOr32(m, v50, int32(_a_F_unlink_span_0), v50)
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
								if v54 != v56 {
									v61 = F_LWLockAcquire(m, v55+int32(1476), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v65+int32(1476))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
												v76 = l0 + v73*int32(20)
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
												if v77 != 0 {
													v81 = v77
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
													return
												} else {
													v78 = F_get_segment_by_index(m, l0, v73)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
														v81 = v80
														v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
														return
													}
												}
											}
										}
									}
								} else {
									v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
									v76 = l0 + v73*int32(20)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									if v77 != 0 {
										v81 = v77
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
										return
									} else {
										v78 = F_get_segment_by_index(m, l0, v73)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
											v81 = v80
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
											return
										}
									}
								}
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v86 = int32(0)
								v89 = base.AtomicRmwOr32(m, v86, int32(_a_F_unlink_span_0), v86)
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+1468))
								if v90 != v92 {
									v97 = F_LWLockAcquire(m, v91+int32(1476), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v101+int32(1476))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
												v112 = l0 + v109*int32(20)
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
												if v113 != 0 {
													v117 = v113
													v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
													v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
													return
												} else {
													v114 = F_get_segment_by_index(m, l0, v109)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
														v117 = v116
														v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
														v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
														return
													}
												}
											}
										}
									}
								} else {
									v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
									v112 = l0 + v109*int32(20)
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
									if v113 != 0 {
										v117 = v113
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
										return
									} else {
										v114 = F_get_segment_by_index(m, l0, v109)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
											v117 = v116
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
											return
										}
									}
								}
							}
						} else {
							v39 = F_get_segment_by_index(m, l0, v34)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
								v42 = v41
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v42+v7&int32(134217727))+4)) = v44
								v46 = v44
								if v46 != 0 {
									v50 = int32(0)
									v53 = base.AtomicRmwOr32(m, v50, int32(_a_F_unlink_span_0), v50)
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
									if v54 != v56 {
										v61 = F_LWLockAcquire(m, v55+int32(1476), int32(0))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v65+int32(1476))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
													v76 = l0 + v73*int32(20)
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
													if v77 != 0 {
														v81 = v77
														v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
														return
													} else {
														v78 = F_get_segment_by_index(m, l0, v73)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
															v81 = v80
															v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
															return
														}
													}
												}
											}
										}
									} else {
										v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
										v76 = l0 + v73*int32(20)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
										if v77 != 0 {
											v81 = v77
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
											return
										} else {
											v78 = F_get_segment_by_index(m, l0, v73)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
												v81 = v80
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
												return
											}
										}
									}
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v86 = int32(0)
									v89 = base.AtomicRmwOr32(m, v86, int32(_a_F_unlink_span_0), v86)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+1468))
									if v90 != v92 {
										v97 = F_LWLockAcquire(m, v91+int32(1476), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v101+int32(1476))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
													v112 = l0 + v109*int32(20)
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
													if v113 != 0 {
														v117 = v113
														v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
														v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
														return
													} else {
														v114 = F_get_segment_by_index(m, l0, v109)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
															v117 = v116
															v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
															v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
															return
														}
													}
												}
											}
										}
									} else {
										v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
										v112 = l0 + v109*int32(20)
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
										if v113 != 0 {
											v117 = v113
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
											return
										} else {
											v114 = F_get_segment_by_index(m, l0, v109)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
												v117 = v116
												v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
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
		} else {
			v34 = int32(base.Ui32(v7) >> (uint(int32(27)) % 32))
			v37 = l0 + v34*int32(20)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
			if v38 != 0 {
				v42 = v38
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v42+v7&int32(134217727))+4)) = v44
				v46 = v44
				if v46 != 0 {
					v50 = int32(0)
					v53 = base.AtomicRmwOr32(m, v50, int32(_a_F_unlink_span_0), v50)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
					if v54 != v56 {
						v61 = F_LWLockAcquire(m, v55+int32(1476), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v65+int32(1476))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
									v76 = l0 + v73*int32(20)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									if v77 != 0 {
										v81 = v77
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
										return
									} else {
										v78 = F_get_segment_by_index(m, l0, v73)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
											v81 = v80
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
											return
										}
									}
								}
							}
						}
					} else {
						v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
						v76 = l0 + v73*int32(20)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
						if v77 != 0 {
							v81 = v77
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
							return
						} else {
							v78 = F_get_segment_by_index(m, l0, v73)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
								v81 = v80
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
								return
							}
						}
					}
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v86 = int32(0)
					v89 = base.AtomicRmwOr32(m, v86, int32(_a_F_unlink_span_0), v86)
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+1468))
					if v90 != v92 {
						v97 = F_LWLockAcquire(m, v91+int32(1476), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v101+int32(1476))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
									v112 = l0 + v109*int32(20)
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
									if v113 != 0 {
										v117 = v113
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
										return
									} else {
										v114 = F_get_segment_by_index(m, l0, v109)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
											v117 = v116
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
											return
										}
									}
								}
							}
						}
					} else {
						v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
						v112 = l0 + v109*int32(20)
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
						if v113 != 0 {
							v117 = v113
							v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
							return
						} else {
							v114 = F_get_segment_by_index(m, l0, v109)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
								v117 = v116
								v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
								return
							}
						}
					}
				}
			} else {
				v39 = F_get_segment_by_index(m, l0, v34)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					v42 = v41
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v42+v7&int32(134217727))+4)) = v44
					v46 = v44
					if v46 != 0 {
						v50 = int32(0)
						v53 = base.AtomicRmwOr32(m, v50, int32(_a_F_unlink_span_0), v50)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
						if v54 != v56 {
							v61 = F_LWLockAcquire(m, v55+int32(1476), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v65+int32(1476))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
										v76 = l0 + v73*int32(20)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
										if v77 != 0 {
											v81 = v77
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
											return
										} else {
											v78 = F_get_segment_by_index(m, l0, v73)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
												v81 = v80
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
												return
											}
										}
									}
								}
							}
						} else {
							v73 = int32(base.Ui32(v46) >> (uint(int32(27)) % 32))
							v76 = l0 + v73*int32(20)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
							if v77 != 0 {
								v81 = v77
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
								return
							} else {
								v78 = F_get_segment_by_index(m, l0, v73)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									v81 = v80
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v81+v46&int32(134217727))+8)) = v83
									return
								}
							}
						}
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v86 = int32(0)
						v89 = base.AtomicRmwOr32(m, v86, int32(_a_F_unlink_span_0), v86)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+1468))
						if v90 != v92 {
							v97 = F_LWLockAcquire(m, v91+int32(1476), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v101+int32(1476))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
										v112 = l0 + v109*int32(20)
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
										if v113 != 0 {
											v117 = v113
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
											return
										} else {
											v114 = F_get_segment_by_index(m, l0, v109)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
												v117 = v116
												v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
												return
											}
										}
									}
								}
							}
						} else {
							v109 = int32(base.Ui32(v85) >> (uint(int32(27)) % 32))
							v112 = l0 + v109*int32(20)
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
							if v113 != 0 {
								v117 = v113
								v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
								return
							} else {
								v114 = F_get_segment_by_index(m, l0, v109)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
									v117 = v116
									v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
									v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v117+v85&int32(134217727)+v119<<(uint(int32(2))%32))+16)) = v123
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
func F_upc_cast_from_ean13(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14315(m, l0, int32(6))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_upc_in(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14257(m, l0, int32(6))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_updateAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	var v11 int32
	_ = v11
	F_updateAclDependenciesWorker(m, l0, l1, l2, l3, int32(97), l4, l5, l6, l7)
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_updateAclDependenciesWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v394 int32
	_ = v394
	var v405 int32
	_ = v405
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v434 int32
	_ = v434
	var v452 int32
	_ = v452
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v524 int32
	_ = v524
	var v553 int32
	_ = v553
	var v564 int64
	_ = v564
	var v565 int64
	_ = v565
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v621 int32
	_ = v621
	var v637 int32
	_ = v637
	var v650 int32
	_ = v650
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v689 int32
	_ = v689
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	v10 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(80)
	m.G0 = v25
	if l5 <= v10 {
		v237 = v10
		v238 = v10
		v240 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v237 < l7 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v38 = v10
	v39 = v10
	v40 = v10
	v42 = v10
	goto L3
L3:
	;
	if v39 < l7 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if l5 <= v87 {
		v237 = v88
		v238 = v89
		v240 = v91
		goto L1
	} else {
		goto L16
	}
L5:
	;
	v52 = int32(2)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l6+v38<<(uint(v52)%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l8+v39<<(uint(v52)%32))))
	if v55 == v59 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v87 = v38
	v88 = v39
	v89 = v40
	v91 = v42
	goto L7
L7:
	;
	goto L4
L8:
	;
	if v82 < l5 {
		v38 = v82
		v39 = v83
		v40 = v84
		v42 = v85
		goto L3
	} else {
		goto L15
	}
L9:
	;
	v61 = int32(1)
	v82 = v38 + v61
	v83 = v39 + v61
	v84 = v40
	v85 = v42
	goto L8
L10:
	;
	goto L11
L11:
	;
	if base.Ui32(v55) < base.Ui32(v59) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v42<<(uint(int32(2))%32)))) = v55
	v70 = int32(1)
	v82 = v38 + v70
	v83 = v39
	v84 = v40
	v85 = v42 + v70
	goto L8
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8+v40<<(uint(int32(2))%32)))) = v59
	v78 = int32(1)
	v82 = v38
	v83 = v39 + v78
	v84 = v40 + v78
	v85 = v42
	goto L8
L15:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	v91 = v85
	goto L7
L16:
	;
	v97 = (l5 - v87) & int32(3)
	if v97 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v160 = l5 + v91 - v87
	if base.Ui32(v87-l5) < base.Ui32(int32(-3)) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v150 = v87
	v151 = v91
	goto L17
L19:
	;
	goto L20
L20:
	;
	v113 = v87
	v114 = v91
	v116 = int32(0)
	goto L21
L21:
	;
	v123 = int32(2)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l6+v113<<(uint(v123)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l6+v114<<(uint(v123)%32)))) = v129
	v131 = int32(1)
	v132 = v114 + v131
	v134 = v113 + v131
	v136 = v116 + v131
	if v136 != v97 {
		v113 = v134
		v114 = v132
		v116 = v136
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v150 = v134
	v151 = v132
	goto L17
L23:
	;
	goto L22
L24:
	;
	v176 = v150
	v177 = v151
	goto L27
L25:
	;
	goto L26
L26:
	;
	v237 = v88
	v238 = v89
	v240 = v160
	goto L1
L27:
	;
	v186 = int32(2)
	v188 = l6 + v177<<(uint(v186)%32)
	v191 = l6 + v176<<(uint(v186)%32)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = v198
	v200 = int32(4)
	v203 = v177 + v200
	if v203 != v160 {
		v176 = v176 + v200
		v177 = v203
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	goto L28
L30:
	;
	v253 = (l7 - v237) & int32(3)
	if v253 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v394 = v238
	goto L32
L32:
	;
	v405 = int32(0)
	if base.B2i32(v240 <= v405)&base.B2i32(v394 <= v405) == v405 {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v316 = l7 + v238 - v237
	if base.Ui32(v237-l7) <= base.Ui32(int32(-4)) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v303 = v237
	v305 = v238
	goto L33
L35:
	;
	goto L36
L36:
	;
	v266 = v237
	v268 = v238
	v269 = int32(0)
	goto L37
L37:
	;
	v279 = int32(2)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l8+v266<<(uint(v279)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l8+v268<<(uint(v279)%32)))) = v285
	v287 = int32(1)
	v288 = v268 + v287
	v290 = v266 + v287
	v292 = v269 + v287
	if v292 != v253 {
		v266 = v290
		v268 = v288
		v269 = v292
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v303 = v290
	v305 = v288
	goto L33
L39:
	;
	goto L38
L40:
	;
	v329 = v303
	v331 = v305
	goto L43
L41:
	;
	goto L42
L42:
	;
	v394 = v316
	goto L32
L43:
	;
	v342 = int32(2)
	v344 = l8 + v331<<(uint(v342)%32)
	v347 = l8 + v329<<(uint(v342)%32)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v350
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v352
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+12)) = v354
	v356 = int32(4)
	v359 = v331 + v356
	if v359 != v316 {
		v329 = v329 + v356
		v331 = v359
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v414 = F_table_open(m, int32(1214), int32(3))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if l6 != 0 {
		goto L105
	} else {
		goto L106
	}
L49:
	;
	return
L50:
	;
	if int32(0) < v394 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v434 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	if int32(0) < v240 {
		goto L93
	} else {
		goto L94
	}
L54:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l8+v434<<(uint(int32(2))%32))))
	if base.B2i32(base.B2i32(l4 != int32(97)) == int32(0))&base.B2i32(v452 == l3) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v581 = v434 + int32(1)
	if v581 != v394 {
		v434 = v581
		goto L54
	} else {
		goto L92
	}
L57:
	;
	v465 = int32(1)
	goto L58
L58:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_updateAclDependenciesWorker_0)) < base.Ui32(v452)) == int32(0))&((v465|base.B2i32(v452 != int32(2200)))&v465) != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	F_shdepLockAndCheckObject(m, int32(1260), v452)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v476 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+11)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v476
	v482 = int32(1)
	if l0 <= int32(3591) {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+64)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = base.I64_extend_i32_u(v452)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = int64(1260)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = base.I64_extend_i32_s(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = base.I64_extend_i32_u(l0)
	v564 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_updateAclDependenciesWorker[0])))
	if v553 != 0 {
		goto L86
	} else {
		goto L87
	}
L62:
	;
	goto L61
L63:
	;
	v553 = int32(0)
	goto L62
L64:
	;
	if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v553 = v482
		goto L62
	} else {
		goto L85
	}
L65:
	;
	if l0 <= int32(2670) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if l0 <= int32(_a_F_updateAclDependenciesWorker_1) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	switch l0 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v553 = v482
		goto L62
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L63
	default:
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v494 = l0 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v494))|base.B2i32(int32(1)<<(uint(v494)%32)&int32(226492515) == int32(0)) != 0 {
		goto L64
	} else {
		goto L73
	}
L71:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v553 = v482
	goto L62
L73:
	;
	v553 = v482
	goto L62
L74:
	;
	if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
		v553 = v482
		goto L62
	} else {
		goto L83
	}
L75:
	;
	v507 = l0 - int32(_a_F_updateAclDependenciesWorker_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v507))|base.B2i32(int32(1)<<(uint(v507)%32)&int32(963) == int32(0)) != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	switch l0 - int32(_a_F_updateAclDependenciesWorker_3) {
	case 0, 1, 2, 3, 4, 59, 60:
		v553 = v482
		goto L62
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L63
	default:
		goto L79
	}
L78:
	;
	v553 = v482
	goto L62
L79:
	;
	if base.Ui32(l0-int32(_a_F_updateAclDependenciesWorker_4)) < base.Ui32(int32(3)) {
		v553 = v482
		goto L62
	} else {
		goto L80
	}
L80:
	;
	v524 = l0 - int32(_a_F_updateAclDependenciesWorker_5)
	if base.Ui32(int32(15)) < base.Ui32(v524) {
		goto L63
	} else {
		goto L81
	}
L81:
	;
	if int32(1)<<(uint(v524)%32)&int32(_a_F_updateAclDependenciesWorker_6) != 0 {
		v553 = v482
		goto L62
	} else {
		goto L82
	}
L82:
	;
	goto L63
L83:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
		goto L63
	} else {
		goto L84
	}
L84:
	;
	v553 = v482
	goto L62
L85:
	;
	goto L63
L86:
	;
	v565 = int64(0)
	goto L88
L87:
	;
	v565 = v564
	goto L88
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v414)+52))
	v572 = F_heap_form_tuple(m, v567, v25+int32(16), v25+int32(8))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L49
	} else {
		goto L89
	}
L89:
	;
	F_CatalogTupleInsert(m, v414, v572)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L49
	} else {
		goto L90
	}
L90:
	;
	F_pfree(m, v572)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L49
	} else {
		goto L91
	}
L91:
	;
	goto L56
L92:
	;
	goto L55
L93:
	;
	v621 = int32(0)
	goto L96
L94:
	;
	goto L95
L95:
	;
	F_relation_close(m, v414, int32(3))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L49
	} else {
		goto L104
	}
L96:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l6+v621<<(uint(int32(2))%32))))
	if base.B2i32(base.B2i32(l4 != int32(97)) == int32(0))&base.B2i32(v637 == l3) != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L95
L98:
	;
	v663 = v621 + int32(1)
	if v663 != v240 {
		v621 = v663
		goto L96
	} else {
		goto L103
	}
L99:
	;
	v650 = int32(1)
	goto L100
L100:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_updateAclDependenciesWorker_0)) < base.Ui32(v637)) == int32(0))&((v650|base.B2i32(v637 != int32(2200)))&v650) != 0 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	F_shdepDropDependency(m, v414, l0, l1, l2, int32(0), int32(1260), v637, l4)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L49
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	goto L97
L104:
	;
	goto L48
L105:
	;
	F_pfree(m, l6)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L49
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if l8 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	F_pfree(m, l8)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L49
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	m.G0 = v25 + int32(80)
	return
L112:
	;
	goto L111
}
func F_update_metainfo_datafile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v9&int32(25) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(144)
	return
L2:
	;
	v15 = F_unlink(m, int32(_a_F_update_metainfo_datafile_0))
	mBase = m.M
	if int32(0) <= v15 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[2]))
	v42 = F_umask(m, v41)
	mBase = m.M
	v45 = F_fopen(m, int32(_a_F_update_metainfo_datafile_4), int32(_a_F_update_metainfo_datafile_5))
	mBase = m.M
	v46 = F_umask(m, v42)
	mBase = m.M
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[1]))
	if v19 == int32(44) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v24 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_update_metainfo_datafile_0)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_1), v6)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1509), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[4]))
	if v115 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+80)) = int32(-1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v96 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L33
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[3]))
	if v57 == int32(0) {
		goto L13
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+80)) = int32(10)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v52 | int32(64)
	goto L17
L21:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v61&int32(1) == int32(0) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+128)) = v57
	v70 = F_pg_fprintf(m, v45, int32(_a_F_update_metainfo_datafile_10), v6+int32(128))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	if int32(0) <= v70 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v76 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v92 = F_fclose(m, v45)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_8), v6+int32(112))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1543), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	goto L1
L33:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_11), v6+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1532), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L1
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[5]))
	if v153 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v119&int32(8) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v115
	v128 = F_pg_fprintf(m, v45, int32(_a_F_update_metainfo_datafile_9), v6+int32(96))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if int32(0) <= v128 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v134 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	if v134 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v150 = F_fclose(m, v45)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_8), v6+int32(80))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1556), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L1
L51:
	;
	v190 = F_fclose(m, v45)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L64
	}
L52:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v157&int32(16) == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v153
	v166 = F_pg_fprintf(m, v45, int32(_a_F_update_metainfo_datafile_7), v6-int32(-64))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	if int32(0) <= v166 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v172 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if v172 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v188 = F_fclose(m, v45)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L63
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_8), v6+int32(48))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1569), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L1
L64:
	;
	v194 = F_rename(m, int32(_a_F_update_metainfo_datafile_4), int32(_a_F_update_metainfo_datafile_0))
	mBase = m.M
	if v194 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v199 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	if v199 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(_a_F_update_metainfo_datafile_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_6), v6+int32(32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1580), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	goto L1
}
func F_utime(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v10 = m.Env.X__syscall_utimensat(m, int32(-100), l0, v2, v2)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v10) {
		*(*int32)(unsafe.Add(mBase, _c_F_utime[0])) = int32(0) - v10
	} else {
	}
	m.G0 = v5 + int32(32)
	return
}
